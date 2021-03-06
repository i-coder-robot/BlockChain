package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type BlockChain struct {
	Blocks          []*Block       `json:"blocks"`
	Difficulty      int            `json:"difficulty"`
	DigMineRewards  float64        `json:"digMineRewards"`
	TransactionPool []*Transaction `json:"TransactionPool"`
	Nodes           []string       `json:"nodes"`
}

type ChainJson struct {
	Length int `json:"length"`
	Chain BlockChain `json:"chain"`
}

func ResolveConflicts(blockChain *BlockChain) bool {
	neighbors := blockChain.Nodes
	var newChain ChainJson
	maxLength := len(blockChain.Blocks)

	for _, node := range neighbors {
		bytes, _, e := HttpGet(node+"/chain")
		fmt.Println(string(bytes))
		if e != nil {
			fmt.Println("请求兄弟节点失败@@@" + e.Error())
		}
		e = json.Unmarshal(bytes, &newChain)
		if e != nil {
			fmt.Println("解析兄弟节点失败@@@" + e.Error())
		}

		length := newChain.Length

		if length > maxLength && newChain.Chain.Validate() {
			maxLength = length
			blockChain.Blocks = newChain.Chain.Blocks
			blockChain.TransactionPool = newChain.Chain.TransactionPool
			blockChain.Nodes = newChain.Chain.Nodes
			return true
		}
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
	block.TimeStamp=time.Now().Unix()
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
		TransactionPool: transactionPool,
		DigMineRewards:  rewards,
		Difficulty:      difficulty,
	}
	return blockChain
}

func (blockChain *BlockChain) String() string {
	return ToString(blockChain)
}
