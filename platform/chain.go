package platform

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/pkg/errors"
	"io/ioutil"
	"math"
	"math/big"
	"net/http"
	"time"

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
	Tokens(string, domain.Network) ([]domain.Token, error)
}

// EthClient holds the necessary values for preforming calls to an EVM network API.
type EthClient struct {
	web3 bind.ContractBackend
}

// tokenList represents the traditional JSON structure used by ERC-20 token lists - tokenlists.org
type tokenList struct {
	Name     string   `json:"name"`
	LogoURI  string   `json:"logoURI"`
	Keywords []string `json:"keywords"`
	Version  struct {
		Major int `json:"major"`
		Minor int `json:"minor"`
		Patch int `json:"patch"`
	} `json:"version"`
	Timestamp time.Time      `json:"timestamp"`
	Tokens    []domain.Token `json:"tokens"`
}

// Tokens scrapes the known list of Token addresses to fetch their name, symbol and (humanised) balance.
// Any token with a balance greater than zero is returned to the domain.Token slice.
func (e EthClient) Tokens(address string, network domain.Network) ([]domain.Token, error) {
	var tokens []domain.Token
	if e.web3 == nil {
		var err error
		e.web3, err = ethclient.Dial(network.Endpoint)
		if err != nil {
			return tokens, err
		}
	}

	var tokenIndex []domain.Token
	for _, l := range network.Tokens.Lists {
		tokenList, err := fetchTokenList(l)
		if err != nil {
			return tokens, errors.Wrap(err, "error fetching token list")
		}
		tokenIndex = append(tokenIndex, tokenList.Tokens...)
	}

	for _, a := range network.Tokens.Additional {
		tokenIndex = append(tokenIndex, domain.Token{Address: a})
	}

	for _, t := range tokenIndex {
		token, _ := contracts.NewToken(common.HexToAddress(t.Address), e.web3)
		balance, _ := token.BalanceOf(common.HexToAddress(address))

		if len(balance.Bits()) > 0 {
			decimals, _ := token.Decimals()
			name, _ := token.Name()
			symbol, _ := token.Symbol()

			tokens = append(tokens, domain.Token{
				Address:  t.Address,
				Symbol:   symbol,
				Name:     name,
				Balance:  bigIntToFloatWithPrecision(balance, int(decimals)),
				Decimals: int(decimals),
			})
		}
	}

	return tokens, nil
}

func fetchTokenList(url string) (tokenList, error) {
	web2 := &http.Client{Timeout: time.Second * 10}

	resp, err := web2.Get(url)
	if err != nil {
		return tokenList{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return tokenList{}, err
	}
	var result tokenList
	if err := json.Unmarshal(body, &result); err != nil {
		return tokenList{}, err
	}

	return result, nil
}

// LiquidityPools scrapes the domain.Market MasterChef contracts to fetch LP & User info
// If the user has a balance/stake in an LP, it is returned to the slice of domain.LiquidityPool
func (e EthClient) LiquidityPools(address string, network domain.Network) ([]domain.LiquidityPool, error) {
	var pools []domain.LiquidityPool
	if e.web3 == nil {
		var err error
		e.web3, err = ethclient.Dial(network.Endpoint)
		if err != nil {
			return pools, err
		}
	}

	for _, m := range network.Markets {
		abi := determineChefABI(m.Name)

		for _, c := range m.Chef {
			marketPools, err := fetchPoolsFromChefContact(
				e.web3, abi, m, common.HexToAddress(address), common.HexToAddress(c))
			if err != nil {
				return []domain.LiquidityPool{}, err
			}

			pools = append(pools, marketPools...)
		}
	}

	return pools, nil
}

func determineChefABI(name string) *bind.MetaData {
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
	client bind.ContractBackend,
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
				Market:  market,
				Address: pool.LpToken.String(),
				Balance: bigIntToFloat(user.Amount),
				Rewards: bigIntToFloat(user.RewardDebt),
				Pair:    pair,
			}
			pools = append(pools, lp)
		}
	}
	return pools, nil
}

// transformTokenPair fetches metadata for the underlying Tokens from LP Token Contract.
// TODO: When Pooled Connections are working fetch Token Symbols.
func transformTokenPair(client bind.ContractBackend, token common.Address) domain.TokenPair {
	// If lpToken doesn't have 'Token0()' assume it's a Token, not a Token Pair (i.e. single sided pool)
	lpToken, _ := contracts.NewTokenPair(token, client)
	token0, err0 := lpToken.Token0(nil)
	token1, err1 := lpToken.Token1(nil)
	if err0 != nil || err1 != nil {
		return domain.TokenPair{
			Token0: domain.Token{Address: token.String()},
			Token1: domain.Token{Address: token.String()},
		}
	}

	return domain.TokenPair{
		Token0: domain.Token{
			Address: token0.String(),
		},
		Token1: domain.Token{
			Address: token1.String(),
		},
	}
}

func bigIntToFloat(amount *big.Int) float64 {
	return bigIntToFloatWithPrecision(amount, 18)
}

func bigIntToFloatWithPrecision(amount *big.Int, decimals int) float64 {
	fAmount, _ := new(big.Float).SetString(amount.String())
	a, _ := new(big.Float).Quo(fAmount, big.NewFloat(math.Pow10(decimals))).Float64()
	return a
}

func (e EthClient) NativeBalance() (string, error) {
	panic("TODO: implement me")
}

func (e EthClient) Balances() ([]string, error) {
	panic("TODO: implement me")
}
