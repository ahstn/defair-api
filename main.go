package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ahstn/defair/internal/contracts"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	address = "<>"
	// eth  = "https://main-light.eth.linkpool.io"
	// eth2 = "https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161"
	avax = "https://api.avax.network/ext/bc/C/rpc"

	joe                = "0x6e84a6216eA6dACC71eE8E6b0a5B7322EEbC0fDd"
	masterchefContract = "0xd6a4f121ca35509af06a0be99093d08462f53052"
)

func main() {
	client, err := ethclient.Dial(avax)
	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	value := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Printf("Balance : %v AVAX \n", value)

	masterchef, err := contracts.NewMasterChef(common.HexToAddress(masterchefContract), client)
	if err != nil {
		log.Fatal(err)
	}

	name, err := masterchef.Joe(nil)
	if err != nil {
		log.Fatalf("Failed to call Joe Masterchef: %v", err)
	}
	fmt.Println("MasterChef 'joe' address:", name)

	l, err := masterchef.PoolLength(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve pool length: %v", err)
	}

	for i := 0; i < int(l.Uint64()); i++ {
		pool, err := masterchef.PoolInfo(nil, big.NewInt(int64(i)))
		if err != nil {
			log.Fatalf("Failed to retrieve pool info: %v", err)
		}
		fmt.Printf("\nJoe LP # %v Token: %s\n", i, pool.LpToken)

		lpToken, err := contracts.NewTokenPair(pool.LpToken, client)
		if err != nil {
			log.Fatalf("Failed to initialise Token Pair: %v", err)
		}
		balance, err := lpToken.BalanceOf(nil, common.HexToAddress(address))
		if err != nil {
			log.Fatalf("Failed to retrieve token balance: %v", err)
		}
		fmt.Printf("Joe LP # %v Balance: %s\n", i, balance)
	}
}
