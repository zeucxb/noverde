package account

import (
	"fmt"
	"noverde/core"
)

func Proccess(accounts core.Data, transactions core.Data) (data map[int]int) {
	data = make(map[int]int)

	for account, transactionValues := range transactions {
		limit := len(transactionValues)
		total := accounts[account][0]

		for key := 0; key < limit; key++ {
			transactionValue := transactionValues[key]

			total += transactionValue

			if total < 0 && transactionValue < 0 {
				total -= 500
			}
		}

		data[account] = total
	}

	return
}

func Show(accounts core.Data, totals map[int]int) {
	for account, value := range accounts {
		if totals[account] == 0 {
			totals[account] = value[0]
		}

		fmt.Printf("%v,%v\n", account, totals[account])
	}
}
