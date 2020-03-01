package BlockChain

import (
	"encoding/json"
	"fmt"
)

type BlockChain struct {
	Blocks []*Block
}

//添加区块到区块链上
func (blockChain *BlockChain) AddBlockToChan(block *Block) *BlockChain {
	blockChain.Blocks = append(blockChain.Blocks, block)
	return blockChain
}

//获取最后一个区块上的hash值
func (blockChain *BlockChain) GetLatestBlockHash() string {
	return blockChain.Blocks[len(blockChain.Blocks)-1].Hash
}

//验证区块链上的数据合法性
func (blockChain *BlockChain) Validate() bool {
	if len(blockChain.Blocks) == 1 {
		return blockChain.Blocks[0].CurrentValidate()
	}

	for i := len(blockChain.Blocks) - 1; i > 0; i-- {
		latest := blockChain.Blocks[i]
		if !latest.CurrentValidate() {
			fmt.Println("数据被别人修改了")
			return false
		}
		pre := blockChain.Blocks[i-1]
		if pre.Hash != latest.PreBlockHash {
			fmt.Println("区块链子断裂")
			return false
		}
	}
	return true
}

//生成老祖宗区块
func Origin(data string, preHash string) *BlockChain {
	originalBlock := NewBlock("我是老祖宗区块", preHash)
	blocks := []*Block{originalBlock}
	blockChain := &BlockChain{
		Blocks: blocks,
	}
	return blockChain
}

func (blockChain *BlockChain) String() {
	bytes, e := json.Marshal(blockChain)
	if e != nil {
		panic("序列化出错拉")
	}
	fmt.Println(string(bytes))
}
