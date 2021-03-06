package actions

import (
	"testing"

	"github.com/ahstn/defair/domain"
	"github.com/ahstn/defair/platform"
	"github.com/golang/mock/gomock"
)

var config = domain.Index{
	Networks: map[string]domain.Network{
		"avalanche": {
			Endpoint: "test.rpc",
			Tokens: domain.TokenSources{
				Additional: []string{"0x123"},
			},
			Markets: []domain.Market{
				{
					Name:  "Trader Joe",
					Token: "Token",
					Chef:  []string{"123"},
				},
			},
		},
		"fantom": {
			Endpoint: "test.rpc",
			Markets: []domain.Market{
				{
					Name:  "SpiritSwap",
					Token: "Token",
					Chef:  []string{"123"},
				},
			},
		},
	},
}

func Test_LiquidityProvider_SingleNetwork(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := platform.NewMockChain(ctrl)
	y := platform.NewMockIndexer(ctrl)
	f := domain.DataFilter{Networks: []string{"avalanche"}}

	y.EXPECT().Read().Return(config, nil).AnyTimes()
	e.EXPECT().
		LiquidityPools(gomock.Any(), config.Networks["avalanche"]).
		Times(1)

	_, err := LiquidityPools("0x123", f, y, e)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_LiquidityProvider_AllNetworks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := platform.NewMockChain(ctrl)
	y := platform.NewMockIndexer(ctrl)
	f := domain.DataFilter{Networks: []string{"all"}}

	y.EXPECT().Read().Return(config, nil).AnyTimes()

	e.EXPECT().
		LiquidityPools(gomock.Any(), config.Networks["avalanche"]).
		Times(1)
	e.EXPECT().
		LiquidityPools(gomock.Any(), config.Networks["harmony"]).
		Times(1)
	e.EXPECT().
		LiquidityPools(gomock.Any(), config.Networks["aurora"]).
		Times(1)

	_, err := LiquidityPools("0x123", f, y, e)
	if err != nil {
		t.Fatal(err)
	}
}
