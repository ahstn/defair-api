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
	LiquidityPools(string, string, []domain.Market) ([]domain.LiquidityPool, error)
}

type EthClient struct{}

// LiquidityPools scrapes the domain.Market MasterChef contracts to fetch LP & User info
// If the user has a balance/stake in an LP, it is returned in the slice of domain.LiquidityPool
//
// So this is a bit of a nightmare. As the members of each "Chef" differ slightly, there's no easy way to make them
//   Generic, at least to my knowledge. Anything that could make "contracts.NewMasterChef" generic, would require
//   a Generic return type. Putting an interface in front of the "Chefs" doesn't work either as their methods all
//   return slightly different structs, i.e. Joe's "PoolInfo" and DFK's "PoolInfo" return different data.
// Maybe future Go Generics help here, or reworking how on-chain data is fetched with ethclient.Client.
// At the very least, there *should* be some similarities in Protocols, i.e. I imagine there are some which haven't
//   tweaked the OG Sushi MasterChef contract.
func (e EthClient) LiquidityPools(address, rpc string, markets []domain.Market) ([]domain.LiquidityPool, error) {
	var pools []domain.LiquidityPool
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return pools, err
	}

	for _, m := range markets {
		caller := DeterminePoolCaller(m.Name)
		marketPools, err := caller(client, m, address)
		if err != nil {
			return []domain.LiquidityPool{}, err
		}

		pools = append(pools, marketPools...)
	}

	return pools, nil
}

// DeterminePoolCaller figures out which MasterChef constructor should be called for a given domain.Market.
// See EthClient.LiquidityPools for context.
// NB: Would it be any simpler if this belonged to domain.Market?
func DeterminePoolCaller(n string) func(client *ethclient.Client, m domain.Market, a string) ([]domain.LiquidityPool, error) {
	switch n {
	case "Defi Kingdoms":
		return defiKingdomPools
	default:
		return traderJoePools
	}
}

// PoolDataToLiquidityPool transforms the on-chain pool data to our common domain.LiquidityPool format
func PoolDataToLiquidityPool(
	client *ethclient.Client,
	market domain.Market,
	amount, reward *big.Int,
	token common.Address,
) (domain.LiquidityPool, error) {
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
		Market: market,
		Address: token.String(),
		Balance: bigIntToFloat(amount),
		Rewards: bigIntToFloat(reward),
		Pair: domain.TokenPair{
			Token0: token0.Hex(),
			Token1: token1.Hex(),
		},
	}, nil
}

func bigIntToFloat(amount *big.Int) float32 {
	fAmount, _ := new(big.Float).SetString(amount.String())
	a, _ := new(big.Float).Quo(fAmount, big.NewFloat(math.Pow10(18))).Float32()
	return a
}

func traderJoePools(client *ethclient.Client, m domain.Market, a string) ([]domain.LiquidityPool, error) {
	var pools []domain.LiquidityPool

	for _, contractAddress := range m.Chef {
		chef, err := contracts.NewMasterChef(common.HexToAddress(contractAddress), client)
		if err != nil {
			return []domain.LiquidityPool{}, err
		}

		l, err := chef.PoolLength(nil)
		if err != nil {
			return []domain.LiquidityPool{}, errors.Wrap(err, "Failed to retrieve pool length")
		}

		for i := 0; i < int(l.Uint64()); i++ {
			pool, err := chef.PoolInfo(nil, big.NewInt(int64(i)))
			if err != nil {
				return pools, errors.Wrap(err, "Failed to retrieve pool information")
			}

			info, err := chef.UserInfo(nil, big.NewInt(int64(i)), common.HexToAddress(a))
			if err != nil {
				return pools, err
			}

			if len(info.Amount.Bits()) > 0 || len(info.RewardDebt.Bits()) > 0 {
				lp, err := PoolDataToLiquidityPool(client, m, info.Amount, info.RewardDebt, pool.LpToken)
				if err != nil {
					return pools, err
				}

				pools = append(pools, lp)
			}
		}
	}

	return pools, nil
}

func defiKingdomPools(client *ethclient.Client, m domain.Market, a string) ([]domain.LiquidityPool, error) {
	var pools []domain.LiquidityPool

	for _, contractAddress := range m.Chef {
		chef, err := contracts.NewGardener(common.HexToAddress(contractAddress), client)
		if err != nil {
			return []domain.LiquidityPool{}, err
		}

		l, err := chef.PoolLength(nil)
		if err != nil {
			return []domain.LiquidityPool{}, errors.Wrap(err, "Failed to retrieve pool length")
		}

		for i := 0; i < int(l.Uint64()); i++ {
			pool, err := chef.PoolInfo(nil, big.NewInt(int64(i)))
			if err != nil {
				return pools, errors.Wrap(err, "Failed to retrieve pool information")
			}

			info, err := chef.UserInfo(nil, big.NewInt(int64(i)), common.HexToAddress(a))
			if err != nil {
				return pools, err
			}

			if len(info.Amount.Bits()) > 0 || len(info.RewardDebt.Bits()) > 0 {
				lp, err := PoolDataToLiquidityPool(client, m, info.Amount, info.RewardDebt, pool.LpToken)
				if err != nil {
					return pools, err
				}

				pools = append(pools, lp)
			}
		}
	}

	return pools, nil
}

func (e EthClient) NativeBalance() (string, error) {
	panic("TODO: implement me")
}

func (e EthClient) Balances() ([]string, error) {
	panic("TODO: implement me")
}
