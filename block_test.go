package main

import (
	"testing"
)

func TestNew(t *testing.T) {
	t1:=&Transaction{
		From:   "",
		To:     "老李",
		Amount: 0,
	}
	newBlock := NewBlock([]*Transaction{t1},"")
	newBlock.String()
}
