package actions

import (
	"github.com/ahstn/defair/domain"
	"github.com/ahstn/defair/platform"
	"github.com/golang/mock/gomock"
	"testing"
)

var (
	stakingConfig = domain.Index{
		Networks: map[string]domain.Network{
			"avalanche": {
				Endpoint: "test.rpc",
				Tokens: domain.TokenSources{
					Additional: []string{"0x123"},
				},
				Markets: []domain.Market{
					{
						Name:  "Trader Joe",
						Token: "0x123",
						Staker: "0x456",
					},
				},
			},
		},
	}
)

func Test_Staking_SingleNetwork(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := platform.NewMockChain(ctrl)
	y := platform.NewMockIndexer(ctrl)
	f := domain.DataFilter{Networks: []string{"avalanche"}}

	y.EXPECT().Read().Return(stakingConfig, nil).AnyTimes()

	for _, market := range stakingConfig.Networks["avalanche"].Markets {
		e.EXPECT().
			TokenInfo(stakingConfig.Networks["avalanche"].Endpoint, gomock.Any(), market.Staker).
			Times(1).
			Return(domain.Token{Balance: 10, TotalSupply: 1000}, nil)

		e.EXPECT().
			TokenInfo(stakingConfig.Networks["avalanche"].Endpoint, market.Staker, market.Token).
			Times(1).
			Return(domain.Token{Balance: 1500}, nil)
	}


	tokens, err := Staking("0x123", f, y, e)
	if err != nil {
		t.Fatal(err)
	}

	if len(tokens) > 1 {
		t.Errorf("Expected Length of Returned Tokens - got = %v , expected = %v", len(tokens), 1)
	}

	if tokens[0].Ratio != 1.5 {
		t.Errorf("Expected Ratio of Returned Token - got = %v , expected = %v", tokens[0].Ratio, 1.5)
	}
}
