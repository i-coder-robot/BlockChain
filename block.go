package BlockChain

import (
	_ "crypto/sha256"
	"encoding/json"
	"fmt"
)

type Block struct {
	Data string
	PreBlockHash string
	Hash string
	Nonce int
}

func NewBlock(data string, preBlockHash string) *Block {
	b := &Block{
		Data:         data,
		PreBlockHash: preBlockHash,
		Hash:         ComputeHash(data,preBlockHash),
	}
	return b
}



func ComputeHash(data string,preBlockHash string) string {
	result :=ProofOfWork(data+preBlockHash)
	return fmt.Sprintf("%x",result)
}

func (block *Block) String() {
	bytes, e := json.Marshal(block)
	if e != nil {
		panic("序列化出错拉")
	}
	fmt.Println(string(bytes))
}

//验证当前区块的数据合法性--校验 数据和preHash的值再次做hash后，是否和传入给你的当前区块的hash值一致
func (block *Block) CurrentValidate() bool {
	computedHash :=ComputeHash(block.Data,block.PreBlockHash)
	return block.Hash == computedHash
}