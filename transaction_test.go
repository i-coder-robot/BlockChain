package BlockChain

import "testing"

func TestNewTransaction(t *testing.T) {

	t1 :=NewTransaction("老王","小姐姐",300.00)
	t2 :=NewTransaction("老张","小姐姐",800.00)
	transactions := []*Transaction{t1,t2}

	b:=&Block{
		Transactions: transactions,
		PreBlockHash: "123456",
		Hash:         "",
		Nonce:        0,
	}
	b.Hash=ComputeHash(b)
	blocks := []*Block{b}
	chain:= NewBlockChain(transactions,blocks,2,50)
	chain.AddBlockToChan(b)
	chain.String()
}
