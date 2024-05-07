package services

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joey1123455/easy_get_coin/storage"
	cryptapi "github.com/joey1123455/go-crypt-api"
)

type StackingContract interface {
	UserTotal(callData *bind.CallOpts, address string) (total *big.Int, err error)
	UserStakeHistory(callData *bind.CallOpts, address string) (res []storage.GameHistoryPayment, err error)
	GeneratePaymentLink(value string) (map[string]interface{}, error)
}

type stakeHistory struct {
	ethClient *ethclient.Client
	contract  *storage.GameHistory
	cryptApi  *cryptapi.Crypt
}

// NewStakingHistory creates a new instance of StakingContract interface implementation
// with the provided Ethereum client and game history contract.
//
// Parameters:
//   - client: An instance of ethclient.Client, the Ethereum client to interact with the blockchain.
//   - contract: An instance of storage.GameHistory, the contract representing the game history storage.
//
// Returns:
//
//	A StakingContract instance representing the staking history contract.
func NewStakingHistory(client *ethclient.Client, contract *storage.GameHistory, cryptClient *cryptapi.Crypt) StackingContract {
	return &stakeHistory{
		ethClient: client,
		contract:  contract,
		cryptApi:  cryptClient,
	}
}

// UserTotal retrieves the total amount staked by a user with the given Ethereum address.
//
// Parameters:
//   - callData: An instance of bind.CallOpts, containing optional parameters for the Ethereum call.
//   - address: The Ethereum address of the user whose staked amount is to be retrieved.
//
// Returns:
//   - res: A *big.Int representing the total amount staked by the user.
//   - err: An error if any occurred during the retrieval process, nil otherwise.
func (g *stakeHistory) UserTotal(callData *bind.CallOpts, address string) (res *big.Int, err error) {
	addressStr := common.HexToAddress(address)
	res, err = g.contract.UserTotal(callData, addressStr)
	return
}

// UserStakeHistory retrieves the staking history of a user with the given Ethereum address.
//
// Parameters:
//   - callData: An instance of bind.CallOpts, containing optional parameters for the Ethereum call.
//   - address: The Ethereum address of the user whose staking history is to be retrieved.
//
// Returns:
//   - res: A slice of storage.GameHistoryPayment representing the staking history of the user.
//   - err: An error if any occurred during the retrieval process, nil otherwise.
func (g *stakeHistory) UserStakeHistory(callData *bind.CallOpts, address string) (res []storage.GameHistoryPayment, err error) {
	addressStr := common.HexToAddress(address)
	res, err = g.contract.UserStakeHistory(callData, addressStr)
	return
}

func (g *stakeHistory) GeneratePaymentLink(value string) (map[string]interface{}, error) {
	_, err := g.cryptApi.GenPaymentAdress()
	if err != nil {
		return nil, err
	}
	paymentAddress, err := g.cryptApi.GenQR("", "250")
	return paymentAddress, err
}
