package platform

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math"
	"math/big"

	"github.com/ahstn/defair/domain"
	"github.com/ahstn/defair/internal/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Chain defines the expected behaviour for a client connecting to any Blockchain
type Chain interface {
	NativeBalance() (string, error)
	Balances() ([]string, error)
	LiquidityPools(string, domain.Network) ([]domain.LiquidityPool, error)
}

type EthClient struct{}

// LiquidityPools scrapes the domain.Market MasterChef contracts to fetch LP & User info
// If the user has a balance/stake in an LP, it is returned to the slice of domain.LiquidityPool
func (e EthClient) LiquidityPools(address string, network domain.Network) ([]domain.LiquidityPool, error) {
	var pools []domain.LiquidityPool
	client, err := ethclient.Dial(network.Endpoint)
	if err != nil {
		return []domain.LiquidityPool{}, err
	}

	for _, m := range network.Markets {
		abi := DetermineChefABI(m.Name)

		for _, c := range m.Chef {
			marketPools, err := fetchPoolsFromChefContact(
				client, abi, m, common.HexToAddress(address), common.HexToAddress(c))
			if err != nil {
				return []domain.LiquidityPool{}, err
			}

			pools = append(pools, marketPools...)
		}
	}

	return pools, nil
}

func DetermineChefABI(name string) *bind.MetaData {
	switch name {
	case "Defi Kingdoms":
		return contracts.DFKChefMetadata
	default:
		return contracts.JoeChefMetadata
	}
}

// fetchPoolsFromChefContact retrieves LPs a wallet is participating in for a single Chef contract.
// This is intended to eventually be called with pipelining/parallelism.
func fetchPoolsFromChefContact(
	client *ethclient.Client,
	abi *bind.MetaData,
	market domain.Market,
	wallet, contract common.Address,
) ([]domain.LiquidityPool, error) {
	chef, err := contracts.NewChefWithABI(contract, abi.ABI, client)
	if err != nil {
		return []domain.LiquidityPool{}, err
	}

	l, err := chef.PoolLength(nil)
	if err != nil {
		return []domain.LiquidityPool{}, err
	}

	var pools []domain.LiquidityPool
	for i := 0; i < int(l.Uint64()); i++ {
		user, err := chef.UserInfo(nil, big.NewInt(int64(i)), wallet)
		if err != nil {
			return []domain.LiquidityPool{}, err
		}

		if len(user.Amount.Bits()) > 0 || len(user.RewardDebt.Bits()) > 0 {
			pool, err := chef.PoolInfo(nil, big.NewInt(int64(i)))
			if err != nil {
				return []domain.LiquidityPool{}, err
			}

			pair := transformTokenPair(client, pool.LpToken)
			lp := domain.LiquidityPool{
				Market: market,
				Address: pool.LpToken.String(),
				Balance: bigIntToFloat(user.Amount),
				Rewards: bigIntToFloat(user.RewardDebt),
				Pair: pair,
			}
			pools = append(pools, lp)
		}
	}
	return pools, nil
}

// transformTokenPair fetches metadata for the underlying Tokens from LP Token Contract
// TODO: Fetch Token Symbols
func transformTokenPair(client *ethclient.Client, token common.Address) (domain.TokenPair) {
	// If lpToken doesn't have 'Token0()' assume it's a Token, not a Token Pair (i.e. single sided pool)
	lpToken, _ := contracts.NewTokenPair(token, client)
	token0, err0 := lpToken.Token0(nil)
	token1, err1 := lpToken.Token1(nil)
	if err0 != nil || err1 != nil {
		return domain.TokenPair{
			Token0: token.String(),
			Token1: token.String(),
		}
	}

	return domain.TokenPair{
		Token0: token0.String(),
		Token1: token1.String(),
	}
}

func bigIntToFloat(amount *big.Int) float32 {
	fAmount, _ := new(big.Float).SetString(amount.String())
	a, _ := new(big.Float).Quo(fAmount, big.NewFloat(math.Pow10(18))).Float32()
	return a
}

func (e EthClient) NativeBalance() (string, error) {
	panic("TODO: implement me")
}

func (e EthClient) Balances() ([]string, error) {
	panic("TODO: implement me")
}
