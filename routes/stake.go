package routes

import (
	"github.com/gin-gonic/gin"
	handler "github.com/joey1123455/easy_get_coin/handlers"
)

type StakeRouteController struct {
	stakeHandler handler.StakeHandler
}

func NewStakeRouteController(stakeHandler handler.StakeHandler) StakeRouteController {
	return StakeRouteController{stakeHandler}
}

// GameDataRoute handles the routes related to game data.
//
// Takes in a gin.RouterGroup as a parameter and does not return anything.
func (r *StakeRouteController) StakeRoute(rg *gin.RouterGroup) {
	router := rg.Group("/stake")

	router.GET("/pay", r.stakeHandler.Stake)
	router.GET("/history/user/:address", r.stakeHandler.UserStakeHistory)
	router.GET("/total/user/:address", r.stakeHandler.UserTotalStake)
}
