package BlockChain

import (
	_ "crypto/sha256"
	"fmt"
)

type Block struct {
	Transactions []*Transaction
	PreBlockHash string
	Hash         string
	Nonce        int
	TimeStamp    int64
}

func NewBlock(data []*Transaction, preBlockHash string) *Block {
	b := &Block{
		Transactions: data,
		PreBlockHash: preBlockHash,
		Hash:         "",
	}
	return b
}

func ComputeHash(block *Block) string {
	data:=string(ToString(block.Transactions))+block.PreBlockHash+string(block.Nonce) + string(block.TimeStamp)
	result := ProofOfWork(data)
	r := fmt.Sprintf("%x", result)
	return r
}

func (block *Block) String() {
	ToString(block)
}

//验证当前区块的数据合法性--校验 数据和preHash的值再次做hash后，是否和传入给你的当前区块的hash值一致
func (block *Block) CurrentValidate() bool {
	computedHash := ComputeHash(block)
	return block.Hash == computedHash
}
