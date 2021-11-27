package platform

import (
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

// LiquidityPools scrapes the domain.Market MasterChef contracts to fetch LP & User info
// If the user has a balance/stake in an LP, it is returned in the slice of domain.LiquidityPool
func (e EthClient) LiquidityPools(address, rpc string, markets []domain.Market) ([]domain.LiquidityPool, error) {
	var pools []domain.LiquidityPool
	for _, m := range markets {
		client, err := ethclient.Dial(rpc)
		if err != nil {
			return pools, err
		}

		chef, err := contracts.NewMasterChef(common.HexToAddress(m.Chef), client)
		if err != nil {
			return pools, err
		}

		l, err := chef.PoolLength(nil)
		if err != nil {
			return pools, errors.Wrap(err, "Failed to retrieve pool length")
		}

		for i := 0; i < int(l.Uint64()); i++ {
			pool, err := chef.PoolInfo(nil, big.NewInt(int64(i)))
			if err != nil {
				return pools, errors.Wrap(err, "Failed to retrieve pool information")
			}
			info, err := chef.UserInfo(nil, big.NewInt(int64(i)), common.HexToAddress(address))
			if err != nil {
				return pools, err
			}

			if len(info.Amount.Bits()) > 0 || len(info.RewardDebt.Bits()) > 0 {
				lp, err := infoToLiquidityPoolStruct(client, info.Amount, info.RewardDebt, pool.LpToken)
				if err != nil {
					return pools, err
				}

				pools = append(pools, lp)
			}
		}
	}

	return pools, nil
}

// TODO: Rename
func infoToLiquidityPoolStruct(client *ethclient.Client, amount, reward *big.Int, token common.Address) (domain.LiquidityPool, error) {
	amountStr, _ := new(big.Float).SetString(amount.String())
	a, _ := new(big.Float).Quo(amountStr, big.NewFloat(math.Pow10(18))).Float32()
	rewardsStr, _ := new(big.Float).SetString(reward.String())
	r, _ := new(big.Float).Quo(rewardsStr, big.NewFloat(math.Pow10(18))).Float32()

	lpToken, err := contracts.NewTokenPair(token, client)
	if err != nil {
		return domain.LiquidityPool{}, errors.Wrap(err, "Failed to initialise Token Pair")
	}

	// If lpToken doesn't have 'Token0()' assume it's a Token, not a Token Pair (i.e. single sided pool)
	token0, err0 := lpToken.Token0(nil)
	token1, err1 := lpToken.Token1(nil)
	if err0 != nil || err1 != nil {
		token0, token1 = token, token
	}

	return domain.LiquidityPool{
		Address: token.String(),
		Balance: a,
		Rewards: r,
		Pair: domain.TokenPair{
			Token0: token0.Hex(),
			Token1: token1.Hex(),
		},
	}, nil
}
