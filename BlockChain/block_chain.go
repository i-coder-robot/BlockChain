package BlockChain

import (
	"encoding/json"
	"fmt"
	BlockChain2 "github.com/i-coder-robot/BlockChain"
)

type BlockChain struct {
	Blocks          []*Block
	Difficulty      int
	DigMineRewards  float64
	transactionPool []*Transaction
	Nodes           []string
}

func ResolveConflicts(blockChain *BlockChain) bool {
	neighbors := blockChain.Nodes
	var newChain *BlockChain
	maxLength := len(blockChain.Blocks)

	for _, node := range neighbors {
		bytes, _, e := BlockChain2.HttpGet(node)
		if e != nil {
			fmt.Println("请求兄弟节点失败@@@"+e.Error())
		}
		e = json.Unmarshal(bytes, newChain)
		if e != nil {
			fmt.Println("解析兄弟节点失败@@@"+e.Error())
		}

		length:= len(newChain.Blocks)

		if length>maxLength && newChain.Validate() {
			maxLength=length
			blockChain = newChain
		}
	}

	if newChain!=nil{
		return true
	}
	return false
}

func (blockChain *BlockChain) RegisterNode(address string) {
	blockChain.Nodes = append(blockChain.Nodes, address)
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

func DigMine(block *Block, difficulty int) *Block {
	mine := ProofOfWorkWithDifficult(block, difficulty)
	block.Hash = mine
	fmt.Printf("恭喜您，挖到矿拉...%s", mine)
	return block
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
func NewBlockChain(transactions []*Transaction, block *Block, difficulty int, rewards float64) *BlockChain {

	blockChain := &BlockChain{
		Blocks:          []*Block{block},
		transactionPool: transactions,
		DigMineRewards:  rewards,
		Difficulty:      difficulty,
	}
	return blockChain
}

func (blockChain *BlockChain) String() string {
	return BlockChain2.ToString(blockChain)
}
