package account

import (
	"fmt"
	"noverde/core"
)

// Proccess accounts and transactions
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

// FormatOutput formats csv output
func FormatOutput(accounts core.Data, totals map[int]int) (output string) {
	for account, value := range accounts {
		if totals[account] == 0 {
			totals[account] = value[0]
		}

		output += fmt.Sprintf("%v,%v\n", account, totals[account])
	}
	return
}
