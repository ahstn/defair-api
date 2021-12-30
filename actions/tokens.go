package actions

import (
	"encoding/json"
	"github.com/ahstn/defair/domain"
	"github.com/ahstn/defair/platform"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"time"
)

// Tokens uses platform.Chain to retrieve a list of token's a Wallet has a positive balance for.
func Tokens(
	address string, filter domain.DataFilter, y platform.Indexer, e platform.Chain,
) ([]domain.Token, error) {
	var tokens []domain.Token

	c, err := y.Read()
	if err != nil {
		return tokens, err
	}

	networks := networkFilter(filter)

	for _, network := range networks {
		var tokenIndex []domain.Token
		for _, l := range c.Networks[network].Tokens.Lists {
			tokenList, err := fetchTokenList(l)
			if err != nil {
				return tokens, errors.Wrap(err, "error fetching token list")
			}
			tokenIndex = append(tokenIndex, tokenList.Tokens...)
		}
		for _, a := range c.Networks[network].Tokens.Additional {
			tokenIndex = append(tokenIndex, domain.Token{Address: a})
		}

		networkTokens, err := e.Balances(c.Networks[network].Endpoint, address, tokenIndex)
		if err != nil {
			return tokens, nil
		}
		tokens = append(tokens, networkTokens...)
	}

	return tokens, nil
}

func fetchTokenList(url string) (domain.TokenList, error) {
	web2 := &http.Client{Timeout: time.Second * 10}

	resp, err := web2.Get(url)
	if err != nil {
		return domain.TokenList{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return domain.TokenList{}, err
	}
	var result domain.TokenList
	if err := json.Unmarshal(body, &result); err != nil {
		return domain.TokenList{}, err
	}

	return result, nil
}
