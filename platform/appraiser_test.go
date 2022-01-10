package platform

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const (
	geckoTokenPrice  = float64(1.35)
	geckoTokenChange = float64(1.1)
	geckoErr         = "asset platform not found"
)

func TestGecko_Appraise(t *testing.T) {
	type args struct {
		network string
		address string
	}

	ts := setupFakeGeckoServer()
	defer ts.Close()

	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr string
	}{
		{
			name: "Happy Path - Successful Response",
			args: args{
				network: "avalanche",
				address: "0x6e84a6216eA6dACC71eE8E6b0a5B7322EEbC0fDd",
			},
			want:    geckoTokenPrice,
			wantErr: "",
		},
		{
			name: "Error - Invalid Network / Asset Platform",
			args: args{
				network: "false-platform",
				address: "0x6e84a6216eA6dACC71eE8E6b0a5B7322EEbC0fDd",
			},
			want:    float64(0),
			wantErr: geckoErr,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := Gecko{baseUrl: fmt.Sprintf("%s/%s", ts.URL, test.args.network)}
			returned, err := g.Appraise(test.args.network, test.args.address)

			if returned != test.want {
				t.Errorf("Appraise() got = %v expected %v", returned, test.want)
			}
			if test.wantErr != "" && err.Error() != test.wantErr {
				t.Errorf("Appraise() got = %s expected %s", err.Error(), test.wantErr)
			}
		})
	}
}

func setupFakeGeckoServer() *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.String(), "avalanche") {
			address := r.URL.Query().Get("contract_addresses")
			resp := map[string]interface{}{
				address: simpleQuote{
					USD:    geckoTokenPrice,
					Change: geckoTokenChange,
				},
			}
			if err := json.NewEncoder(w).Encode(resp); err != nil {
				panic(err)
			}
		} else {
			if err := json.NewEncoder(w).Encode(geckoError{Error: geckoErr}); err != nil {
				panic(err)
			}
		}
	}))
	return ts
}
