package actions

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/ahstn/defair/domain"
	"github.com/ahstn/defair/platform"
	"github.com/golang/mock/gomock"
)

var fakeTokenList = domain.TokenList{
	Name: "Test",
	Tokens: []domain.Token{
		{
			Address: "0x456",
			Name:    "Test Token",
			Symbol:  "TEST",
		},
	},
}

func Test_Tokens_SingleNetwork(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := platform.NewMockChain(ctrl)
	y := platform.NewMockIndexer(ctrl)
	f := domain.DataFilter{Networks: []string{"avalanche"}}

	// Setup Fake HTTP Server and alter Indexer's config to include the URL as a token list
	ts := setupFakeTokenListServer()
	defer ts.Close()
	if entry, ok := config.Networks["avalanche"]; ok {
		entry.Tokens.Lists = []string{ts.URL}
		config.Networks["avalanche"] = entry
	}

	// Tokens should include addresses returned by platform.Indexer & HTTP Token List
	tokens := fakeTokenList.Tokens
	tokens = append(tokens, domain.Token{Address: config.Networks["avalanche"].Tokens.Additional[0]})

	y.EXPECT().Read().Return(config, nil).AnyTimes()
	e.EXPECT().
		Balances(gomock.Any(), gomock.Any(), tokens).
		Times(1)

	_, err := Tokens("0x123", f, y, e)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFetchTokenList(t *testing.T) {
	ts := setupFakeTokenListServer()
	defer ts.Close()

	response, err := fetchTokenList(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(fakeTokenList, response) {
		t.Errorf("Read() got = %+v, expected = %+v", response, fakeTokenList)
	}
}

func setupFakeTokenListServer() *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(fakeTokenList)
		if err != nil {
			panic(err)
		}
	}))
	return ts
}
