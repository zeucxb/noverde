package main

import (
	"noverde/account"
	"noverde/core"
)

func main() {
	accountsPath, transactionsPath := core.GetArgs()

	accounts := core.Parse(accountsPath)
	transactions := core.Parse(transactionsPath)

	totals := account.Proccess(accounts, transactions)

	account.Show(accounts, totals)
}
