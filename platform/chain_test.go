package platform

import (
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

// TestBalances mocks ethclient to return pre-determined values for a Token contract and it's methods:
//   "balanceOf", "decimals", "symbol" & "name".
// In this instance, we're calling a Token Contract Address to check the balance of "walletAddress". We do this to
//   validate the correct methods are called and with the correct data being returned.
func TestBalances(t *testing.T) {
	walletAddressString := "0x456"
	walletAddress := common.HexToAddress(walletAddressString)
	tokenAddressString := abcToken.Address
	tokenAddress := common.HexToAddress(tokenAddressString)
	tokenIndex := []domain.Token{{Address: abcToken.Address}}

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

	tokens, err := e.Balances("rpc://", walletAddressString, tokenIndex)
	if err != nil {
		t.Fatal(err)
	}

	expected := []domain.Token{abcToken}
	if !reflect.DeepEqual(expected, tokens) {
		t.Errorf("Balances() got = %+v , expected = %+v", tokens, expected)
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
