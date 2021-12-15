package platform

import (
	"encoding/json"
	"fmt"
	"github.com/ahstn/defair/domain"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestFetchTokenList(t *testing.T) {
	data := TokenList{
		Name:     "Test",
		Tokens:    []domain.Token{
			{
				Address: "0x123",
				Name: "Test Token",
				Symbol: "TEST",
			},
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(data)
	}))
	defer ts.Close()

	response, err := fetchTokenList(ts.URL)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(response)

	if !reflect.DeepEqual(data, response) {
		t.Errorf("Read() got = %+v, expected = %+v", response, data)
	}
}
