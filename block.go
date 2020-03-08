package main

import (
	_ "crypto/sha256"
	"fmt"
)

type Block struct {
	Transactions []*Transaction `json:"transactions"`
	PreBlockHash string         `json:"preBlockHash"`
	Hash string                 `json:"hash"`
	Nonce int                   `json:"nonce"`
	//1393474453543 <--> 2020-03-05 21:00:05
	TimeStamp int64 `json:"timeStamp"`
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
	data := string(ToString(block.Transactions)) + block.PreBlockHash + string(block.Nonce) + string(block.TimeStamp)
	result := ProofOfWork(data)
	return fmt.Sprintf("%x",result)
}

func (block *Block) String() string {
	return ToString(block)
}

//验证当前区块的数据合法性--校验 数据和preHash的值再次做hash后，是否和传入给你的当前区块的hash值一致
func (block *Block) CurrentValidate() bool {
	computedHash :=ProofOfWorkWithDifficult(block,blockChain.Difficulty)
	return block.Hash == computedHash
}