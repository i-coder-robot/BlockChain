package BlockChain

import (
	"testing"
)

func TestNew(t *testing.T) {
	t1:=&Transaction{
		From:   "老王",
		To:     "小姐姐",
		Amount: 0.68,
	}
	newBlock := NewBlock([]*Transaction{t1},"123456")
	newBlock.String()
}
