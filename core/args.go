package core

import (
	"log"
	"os"
)

func checkArgs() {
	if len(os.Args) != 3 {
		log.Fatal("Argumentos inválidos!")
	}
}

func GetArgs() (accountsPath string, transactionsPath string) {
	checkArgs()

	accountsPath = os.Args[1]
	transactionsPath = os.Args[2]

	return
}
