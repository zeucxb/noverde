package account

import (
	"noverde/core"
	"testing"

	. "github.com/franela/goblin"
)

func TestProccess(t *testing.T) {
	g := Goblin(t)
	g.Describe("Account", func() {
		g.It("Proccess accounts and transactions", func() {
			accounts := make(core.Data)
			transactions := make(core.Data)

			accounts[123] = []int{10000}
			accounts[234] = []int{15000}
			accounts[345] = []int{0}
			accounts[456] = []int{-30}
			accounts[567] = []int{-30}

			transactions[123] = []int{5000, -200, -30000, 9000, -100}
			transactions[234] = []int{-15000, -200, -1}
			transactions[567] = []int{-1000, -10, 30000}

			data := Proccess(accounts, transactions)

			g.Assert(data[123]).Equal(-7300)
			g.Assert(data[234]).Equal(-1201)
			g.Assert(data[345]).Equal(0)
			g.Assert(data[456]).Equal(0) // This is the transactions result
			g.Assert(data[567]).Equal(27960)
		})
	})
}
