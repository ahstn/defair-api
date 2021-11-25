package platform

import (
	"fmt"
	"math"
	"math/big"

	"github.com/ahstn/defair/domain"
	"github.com/ahstn/defair/internal/contracts"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common"
)

// Chain defines the expected behaviour for a client connecting to any Blockchain
type Chain interface {
	NativeBalance() (string, error)
	Balances() ([]string, error)
	LiquidityPools()
}

type EthClient struct{}

func (e EthClient) LiquidityPools(address, rpc string, markets []domain.Market) error {
	for _, m := range markets {
		client, err := ethclient.Dial(rpc)
		if err != nil {
			return err
		}

		chef, err := contracts.NewMasterChef(common.HexToAddress(m.Chef), client)
		if err != nil {
			return err
		}

		l, err := chef.PoolLength(nil)
		if err != nil {
			return errors.Wrap(err, "Failed to retrieve pool length")
		}

		for i := 0; i < int(l.Uint64()); i++ {
			pool, err := chef.PoolInfo(nil, big.NewInt(int64(i)))
			if err != nil {
				return errors.Wrap(err, "Failed to retrieve pool information")
			}

			info, err := chef.UserInfo(nil, big.NewInt(int64(i)), common.HexToAddress(address))
			if err != nil {
				return err
			}

			if len(info.Amount.Bits()) > 0 || len(info.RewardDebt.Bits()) > 0 {
				fmt.Printf("\nLP #%v Token: %s\n", i, pool.LpToken)

				amountStr, _ := new(big.Float).SetString(info.Amount.String())
				amount := new(big.Float).Quo(amountStr, big.NewFloat((math.Pow10(18))))
				rewardsStr, _ := new(big.Float).SetString(info.RewardDebt.String())
				rewards := new(big.Float).Quo(rewardsStr, big.NewFloat(math.Pow10(18)))
				fmt.Printf("LP #%v Staked Balance: %v , Rewards: %v \n", i, amount, rewards)

				lpToken, err := contracts.NewTokenPair(pool.LpToken, client)
				if err != nil {
					return errors.Wrap(err, "Failed to initialise Token Pair")
				}

				balance, err := lpToken.BalanceOf(nil, common.HexToAddress(address))
				if err != nil {
					return errors.Wrap(err, "Failed to retrieve token balance")
				}
				fmt.Printf("LP #%v Wallet Balance: %s\n", i, balance)
			}
		}
	}

	return nil
}
