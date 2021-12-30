package actions

import (
	"github.com/ahstn/defair/domain"
	"github.com/ahstn/defair/platform"
)

// Staking uses platform.Chain to retrieve a list of token's a Wallet has a positive balance for.
func Staking(
	address string, filter domain.DataFilter, y platform.Indexer, e platform.Chain,
) ([]domain.Token, error) {
	var tokens []domain.Token

	c, err := y.Read()
	if err != nil {
		return tokens, err
	}

	networks := networkFilter(filter)

	for _, network := range networks {
		for _, m := range c.Networks[network].Markets {
			if m.Staker != "" {
				stakeToken, err := e.TokenInfo(c.Networks[network].Endpoint, address, m.Staker)
				if err != nil {
					return tokens, err
				}

				if stakeToken.Balance > 0 {
					govToken, err := e.TokenInfo(c.Networks[network].Endpoint, m.Staker, m.Token)
					if err != nil {
						return tokens, err
					}

					stakeToken.Ratio = govToken.Balance / stakeToken.TotalSupply
					tokens = append(tokens, stakeToken)
				}
			}
		}
	}

	return tokens, nil
}
