package BlockChain

import (
	_ "crypto/sha256"
	"fmt"
	BlockChain2 "github.com/i-coder-robot/BlockChain"
)

type Block struct {
	Transactions []*Transaction
	PreBlockHash string
	Hash string
	Nonce int
	//1393474453543 <--> 2020-03-05 21:00:05
	TimeStamp int64
}

func NewBlock(transactions []*Transaction, preBlockHash string) *Block {
	b := &Block{
		Transactions: transactions,
		PreBlockHash: preBlockHash,
		Hash:         "",
	}
	return b
}



func ComputeHash(block *Block) string {
	data := string(BlockChain2.ToString(block.Transactions)) + block.PreBlockHash + string(block.Nonce) + string(block.TimeStamp)
	result := ProofOfWork(data)
	return fmt.Sprintf("%x",result)
}

func (block *Block) String() string {
	return BlockChain2.ToString(block)
}

//验证当前区块的数据合法性--校验 数据和preHash的值再次做hash后，是否和传入给你的当前区块的hash值一致
func (block *Block) CurrentValidate() bool {
	computedHash :=ComputeHash(block)
	return block.Hash == computedHash
}