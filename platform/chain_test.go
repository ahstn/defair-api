package platform

import (
	"encoding/json"
	"fmt"
	"github.com/ahstn/defair/domain"
	"github.com/ahstn/defair/internal/contracts"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethmath "github.com/ethereum/go-ethereum/common/math"
	"github.com/golang/mock/gomock"
	"math"
	"math/big"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

var (
	abcToken = domain.Token{
		Address: "0x123",
		Balance: 2,
		Decimals: 18,
		Name: "Test Token",
		Symbol: "ABC",
	}
)

// TestTokens mocks ethclient to return pre-determined values for a Token contract and it's methods:
//   "balanceOf", "decimals", "symbol" & "name".
// In this instance, we're calling a Token Contract Address to check the balance of "walletAddress". We do this to
//   validate the correct methods are called and with the correct data being returned.
func TestTokens(t *testing.T) {
	walletAddressString := "0x456"
	walletAddress := common.HexToAddress(walletAddressString)
	tokenAddressString := "0x123"
	tokenAddress := common.HexToAddress(tokenAddressString)
	network := domain.Network{
		Endpoint:  "rpc://",
		Tokens:    domain.TokenSources{
			Additional: []string{
				tokenAddressString,
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	c := NewMockContractBackend(ctrl)
	e := EthClient{web3: c}

	parsedABI, _ := abi.JSON(strings.NewReader(contracts.TokenMetaData.ABI))

	callSignature, _ := parsedABI.Pack("balanceOf", walletAddress)
	returnValue := ethmath.U256Bytes(
		big.NewInt(0).Mul(
			big.NewInt(int64(abcToken.Balance)), big.NewInt(int64(math.Pow10(abcToken.Decimals))),
		),
	)
	c.EXPECT().
		CallContract(gomock.Any(), ethereum.CallMsg{
			From: common.HexToAddress("0x0"),
			To: &tokenAddress,
			Data: callSignature,
		}, gomock.Any()).
		Return(returnValue, nil).AnyTimes()

	callSignature, _ = parsedABI.Pack("decimals")
	returnValue = ethmath.U256Bytes(big.NewInt(int64(abcToken.Decimals)))
	c.EXPECT().
		CallContract(gomock.Any(), ethereum.CallMsg{
			From: common.HexToAddress("0x0"),
			To: &tokenAddress,
			Data: callSignature,
		}, gomock.Any()).
		Return(returnValue, nil).AnyTimes()

	callSignature, _ = parsedABI.Pack("name")
	returnValue, _ = packString(abcToken.Name)
	c.EXPECT().
		CallContract(gomock.Any(), ethereum.CallMsg{
			From: common.HexToAddress("0x0"),
			To:   &tokenAddress,
			Data: callSignature,
		}, gomock.Any()).
		Return(returnValue, nil).AnyTimes()

	callSignature, _ = parsedABI.Pack("symbol")
	returnValue, _ = packString(abcToken.Symbol)
	c.EXPECT().
		CallContract(gomock.Any(), ethereum.CallMsg{
			From: common.HexToAddress("0x0"),
			To:   &tokenAddress,
			Data: callSignature,
		}, gomock.Any()).
		Return(returnValue, nil).AnyTimes()

	tokens, err := e.Tokens(walletAddressString, network)
	if err != nil {
		t.Fatal(err)
	}

	expected := []domain.Token{abcToken}
	if !reflect.DeepEqual(expected, tokens) {
		t.Errorf("Tokens() got = %+v , expected = %+v", tokens, expected)
	}
}

func TestFetchTokenList(t *testing.T) {
	data := tokenList{
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

func packString(val string) ([]byte, error) {
	inDef := fmt.Sprintf(`[{ "name" : "method", "type": "function", "inputs": %s}]`, `[{"type": "string"}]`)
	inAbi, err := abi.JSON(strings.NewReader(inDef))
	if err != nil {
		return []byte(""), err
	}
	var packed []byte
	packed, err = inAbi.Pack("method", val)
	if err != nil {
		return []byte(""), err
	}

	return packed[4:], nil
}
