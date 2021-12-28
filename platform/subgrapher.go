package platform

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	"log"
	"strconv"
)

type Subgrapher interface {
}

// Subgraph is a simple low-level implementation using the following library:
// https://github.com/machinebox/graphql.
type Subgraph struct {
}

// Pools uses the `LiquidityPoolsSnapshots` to calculate a user's overall position in LPs.
// Each Snapshot is how much a user added at a given block, so we add them together to determine the total position.
// Token amount is calculated by: (LP Token Balance / LP Token Total Supply) * Underlying Token X Reserve Supply.
func Pools() {
	client := graphql.NewClient("http://graph3.defikingdoms.com/subgraphs/name/defikingdoms/dex")

	req := graphql.NewRequest(snapshotsQuery)

	req.Var("addr", "$ADDR")
	req.Header.Set("Cache-Control", "no-cache")

	var respData SnapshotsResponse
	if err := client.Run(context.Background(), req, &respData); err != nil {
		log.Fatal(err)
	}

	balances := make(map[string]Positions)
	for _, snapshot := range respData.Snapshots {
		bal, _ := strconv.ParseFloat(snapshot.Balance, 64)
		supply, _ := strconv.ParseFloat(snapshot.Pool.Pair.TotalSupply, 64)
		reserve0, _  := strconv.ParseFloat(snapshot.Pool.Pair.Reserve0, 64)
		reserve1, _  := strconv.ParseFloat(snapshot.Pool.Pair.Reserve1, 64)
		share := bal / supply

		snapshot.Pool.Balance = balances[snapshot.Pool.ID].Balance + bal

		// Determine the token amounts by assigning the current amount, or if we already have a balance,
		//   add to the already existing amounts.
		snapshot.Pool.Pair.Token0.Amount = share * reserve0
		snapshot.Pool.Pair.Token1.Amount = share * reserve1
		if balances[snapshot.Pool.ID].Balance > 0 {
			snapshot.Pool.Pair.Token0.Amount = balances[snapshot.Pool.ID].Pair.Token0.Amount + (share * reserve0)
			snapshot.Pool.Pair.Token1.Amount = balances[snapshot.Pool.ID].Pair.Token1.Amount + (share * reserve1)
		}

		balances[snapshot.Pool.ID] = snapshot.Pool
	}
	fmt.Printf("%+v\n", balances)
}
