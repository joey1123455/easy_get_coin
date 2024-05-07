package handler

import (
	"context"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gin-gonic/gin"
	"github.com/joey1123455/easy_get_coin/services"
	"github.com/joey1123455/easy_get_coin/storage"
	"github.com/joey1123455/easy_get_coin/utils"
)

type StakeHandler struct {
	services        services.StackingContract
	ctx             *context.Context
	TransactOpts    *bind.TransactOpts
	CallOpts        *bind.CallOpts
	Cache           *utils.Cache
	contractAddress string
}

// NewStakingHandler creates a new StakeHandler instance.
//
// Parameters:
//
//		service: services.StackingContract
//		ctx_: *context.Context
//		tans: *bind.TransactOpts
//		call: *bind.CallOpts
//		cache: *utils.Cache
//	 	contractAddress: string
//
// Return Type:
//
//	*StakeHandler
func NewStakingHandler(service services.StackingContract, ctx_ *context.Context, tans *bind.TransactOpts, call *bind.CallOpts, cache *utils.Cache, contractAdd string) *StakeHandler {
	return &StakeHandler{
		services:        service,
		ctx:             ctx_,
		TransactOpts:    tans,
		CallOpts:        call,
		Cache:           cache,
		contractAddress: contractAdd,
	}
}

// UserTotalStake godoc
// @Summary      Show game history
// @Description  handles the retrieval total stake for a users wallet..
// @Tags         staking
// @Produce      json
// @Param        address   path      string  true  "Wallet Address"
// @Success      200  {object}  handler.GameHistoryResOk
// @Failure      400  {object}  handler.GameHistoryResFail
// @Failure      404  {object}  handler.GameHistoryResFail
// @Failure      500  {object}  handler.GameHistoryResFail
// @Router       /stake/total/user/{address} [get]
func (g *StakeHandler) UserTotalStake(ctx *gin.Context) {
	address := ctx.Param("address")

	res, err := g.services.UserTotal(g.CallOpts, address)
	if err != nil {
		log.Println("while getting game data: ", err.Error())
		response := GameHistoryResFail{
			Status:  "fail",
			Message: err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// Stake godoc
// @Summary      Generate payment button and QRCode
// @Description  Generates the qr code and payment link for user stake.
// @Tags         staking
// @Produce      json
// @Param        value   query      string  true  "Transaction Value"
// @Success      200  {object}  handler.GameHistoryResOk
// @Failure      400  {object}  handler.GameHistoryResFail
// @Failure      404  {object}  handler.GameHistoryResFail
// @Failure      500  {object}  handler.GameHistoryResFail
// @Router       /stake/pay [get]
func (g *StakeHandler) Stake(ctx *gin.Context) {
	value := ctx.Query("value")

	res, err := g.services.GeneratePaymentLink(value)
	if err != nil {
		log.Println("while generating payment qr code: ", err.Error())
		response := GameHistoryResFail{
			Status:  "fail",
			Message: err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// UserStakeHistory godoc
// @Summary      Show user game history
// @Description  handles the retrieval of stake history for a given wallet. It paginates the results based on the page and pageSize query parameters.
// @Tags         staking
// @Produce      json
// @Param        address   path      string  true  "Wallet Address"
// @Param        page  query     string     false  "Page number"
// @Param        pageSize  query     string     false  "Page size"
// @Success      200  {object}  handler.GameHistoryResOk
// @Failure      400  {object}  handler.GameHistoryResFail
// @Failure      404  {object}  handler.GameHistoryResFail
// @Failure      500  {object}  handler.GameHistoryResFail
// @Router       /stake/history/user/{address} [get]
func (g *StakeHandler) UserStakeHistory(ctx *gin.Context) {
	var res []storage.GameHistoryPayment
	address := ctx.Param("address")

	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		log.Println("Invalid page number")
		response := GameHistoryResFail{
			Status:  "fail",
			Message: "Invalid page number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	pageSize, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))
	if err != nil || pageSize < 1 {
		log.Println("Invalid page size")
		response := GameHistoryResFail{
			Status:  "fail",
			Message: "Invalid page size",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if _, found := g.Cache.Get(address); !found {
		res, err = g.services.UserStakeHistory(g.CallOpts, address)
		if err != nil {
			log.Println("while getting game data: ", err.Error())
			response := GameHistoryResFail{
				Status:  "fail",
				Message: err.Error(),
			}
			ctx.JSON(http.StatusBadRequest, response)
			return
		}

		sort.SliceStable(res, func(i, j int) bool {
			return utils.ComparePtrFieldsDesc(&res[i], &res[j])
		})

		g.Cache.Set(address, res, 6*time.Minute)
	}

	cachedData, _ := g.Cache.Get(address)
	res = cachedData.([]storage.GameHistoryPayment)

	if len(res) == 0 {
		response := GameHistoryResOk{
			Status: "failed no game data for provided gid",
			Page:   []storage.GameHistoryGameSession{},
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	startIndex := (page - 1) * pageSize
	endIndex := page * pageSize
	if startIndex >= len(res) {
		response := GameHistoryResOk{
			Status: "failed no new page",
			Page:   []storage.GameHistoryGameSession{},
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	if endIndex > len(res) {
		endIndex = len(res)
	}
	response := GameHistoryResOk{
		Status: "success",
		Page:   res[startIndex:endIndex],
	}
	ctx.JSON(http.StatusOK, response)
}
