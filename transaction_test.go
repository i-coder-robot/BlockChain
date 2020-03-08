package main

import (
	"fmt"
	"testing"
)

func TestNewTransaction(t *testing.T) {
	t0:=&Transaction{
		From:   "",
		To:     "老李",
		Amount: 2.0,
	}
	//r:=t1.String()
	//fmt.Println(r)

	t1:=&Transaction{
		From:   "小明",
		To:     "小红",
		Amount: 2.0,
	}
	t2:=&Transaction{
		From:   "小红",
		To:     "老王",
		Amount: 1.8,
	}
	transactions :=[]*Transaction{t0,t1,t2}
	b:=&BlockChain.Block{
		Transactions:transactions,
		PreBlockHash: "123456",
	}
	b.Hash= BlockChain.ComputeHash(b)
	BlockChain.DigMine(b,2)
	chain:= BlockChain.NewBlockChain(transactions,b,2,50.0)
	r:=chain.String()

	fmt.Println(r)
}
