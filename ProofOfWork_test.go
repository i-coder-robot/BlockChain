package BlockChain

import (
	"fmt"
	"testing"
)

func TestProofOfWork(t *testing.T) {
	fmt.Printf("%x",ProofOfWork("block_chan_1"))
	fmt.Printf("\n%x",ProofOfWork("block_chan_2"))
	fmt.Printf("\n%x",ProofOfWork("区块链"))
}

func TestProofOfWorkWithDifficult(t *testing.T) {
	ProofOfWorkWithDifficult("block_chan_1",2)
	//ProofOfWorkWithDifficult("block_chan_2",2)
}