package BlockChain

import (
	"fmt"
	"testing"
	"time"
)

func TestProofOfWork(t *testing.T) {
	fmt.Printf("%x", ProofOfWork("block_chan_1"))
	fmt.Printf("\n%x", ProofOfWork("block_chan_2"))
	fmt.Printf("\n%x", ProofOfWork("区块链"))
}

func TestProofOfWorkWithDifficult(t *testing.T) {
	t1 := &Transaction{
		From:   "小明",
		To:     "小姐姐",
		Amount: 2.101,
	}
	b := &Block{
		Transactions: []*Transaction{t1},
		PreBlockHash: "123456",
		Hash:         "",
		Nonce:        0,
		TimeStamp:    time.Now().Unix(),
	}
	b.Hash = ComputeHash(b)
	ProofOfWorkWithDifficult(b, 2)
}
