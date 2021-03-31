package stats

import (
	"github.com/firuz01/bank/pkg/bank/types"
	"fmt"
	
)

func ExampleAvg()  {
	payments := []types.Payment {
		{
			ID: 1,
			Amount: 60,
			Category: "shop",
		},{
			ID: 2,
			Amount: 80,
			Category: "shop",

		},
	}
	fmt.Println(Avg(payments))
	// Output: 70
}