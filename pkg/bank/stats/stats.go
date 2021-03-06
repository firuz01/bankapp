package stats

import (
	"github.com/firuz01/bank/pkg/bank/types"

)

func Avg(payments []types.Payment) types.Money {
	numPayments := types.Money(len(payments))
	sumPayments := types.Money(0)
	for _, payment := range payments {
		moneyAmount := payment.Amount
		sumPayments += moneyAmount
	}
	return sumPayments/numPayments
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {

		sumPayments := types.Money(0)
		for _, payment := range payments {
			if payment.Category != category {
				continue
			}
			sumPayments += payment.Amount
		}
		return sumPayments
}