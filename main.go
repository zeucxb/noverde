package main

import (
	"fmt"
	"noverde/account"
	"noverde/core"
)

func main() {
	accountsPath, transactionsPath := core.GetArgs()

	accounts := core.ParseAndRead(accountsPath)
	transactions := core.ParseAndRead(transactionsPath)

	totals := account.Proccess(accounts, transactions)

	output := account.FormatOutput(accounts, totals)

	fmt.Println(output)
}
