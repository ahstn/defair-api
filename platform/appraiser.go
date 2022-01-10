package platform

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Appraiser interface {
	Appraise(string, string) (float64, error)
}

// Gecko is an implementation of Appraiser for https://coingecko.com
type Gecko struct {
	client  *http.Client
	baseUrl string
}

type simpleQuote struct {
	USD    float64 `json:"usd"`
	Change float64 `json:"usd_24h_change"`
}

type geckoError struct {
	Error string `json:"error"`
}

const (
	baseUrl = "https://api.coingecko.com/api/v3/simple/token_price/%s"
)

// Appraise returns a price quote for a single asset given it's parent network & contract address
func (g *Gecko) Appraise(network, address string) (float64, error) {
	if g.client == nil {
		g.client = &http.Client{Timeout: time.Second * 10}
	}
	if g.baseUrl == "" {
		g.baseUrl = fmt.Sprintf(baseUrl, network)
	}

	req, err := http.NewRequest(http.MethodGet, g.baseUrl, nil)
	if err != nil {
		return 0, err
	}

	q := req.URL.Query()
	q.Add("contract_addresses", address)
	q.Add("vs_currencies", "usd")
	req.URL.RawQuery = q.Encode()

	resp, err := g.client.Do(req)
	if err != nil {
		return 0, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	var result map[string]simpleQuote
	if err := json.Unmarshal(body, &result); err != nil {
		var errResp geckoError
		if err := json.Unmarshal(body, &errResp); err != nil {
			return 0, err
		}

		return 0, errors.New(errResp.Error)
	}

	return result[address].USD, nil
}
