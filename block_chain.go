package BlockChain

import (
	"fmt"
)

type BlockChain struct {
	Blocks          []*Block
	Difficulty      int
	DigMineRewards  int
	TransactionPool []*Transaction
}

//生成老祖宗区块
func NewBlockChain(transactions []*Transaction, blocks []*Block,difficulty,rewards int) *BlockChain {

	blockChain := &BlockChain{
		Blocks: blocks,
		TransactionPool:transactions,
		Difficulty:difficulty,
		DigMineRewards:rewards,
	}
	return blockChain
}


//添加区块到区块链上
func (blockChain *BlockChain) AddBlockToChan(block *Block) *BlockChain {
	DigMine(block, blockChain.Difficulty)
	blockChain.Blocks = append(blockChain.Blocks, block)
	return blockChain
}

func DigMine(block *Block,difficulty int) *Block {
	mine:=ProofOfWorkWithDifficult(block,difficulty)
	block.Hash=mine
	fmt.Printf("恭喜你，挖到矿了...%s\n",mine)
	return block
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


func (blockChain *BlockChain) String() {
	ToString(blockChain)
}
