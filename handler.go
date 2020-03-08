package main

import (
	"github.com/gin-gonic/gin"
	"github.com/i-coder-robot/BlockChain/BlockChain"
	"net/http"
	"strings"
	"time"
)

var blockChain *BlockChain.BlockChain
const (
	rewards = 50.0
	difficulty = 4
)
func Init() {
	transaction:=&BlockChain.Transaction{
		From:   "",
		To:     "欢喜哥",
		Amount: rewards,
	}
	transactions := []*BlockChain.Transaction{transaction}
	block := &BlockChain.Block{
		Transactions: transactions,
		PreBlockHash: "",
		Hash:         "",
		Nonce:        0,
		TimeStamp:    time.Now().Unix(),
	}
	block.Hash = BlockChain.ComputeHash(block)
	blockChain = BlockChain.NewBlockChain(transactions, block, difficulty, rewards)
}

func MineHandler(c *gin.Context) {

	r := blockChain.Validate()
	if !r{
		panic("该区块链校验失败")
	}
	currentBlock:=blockChain.Blocks[len(blockChain.Blocks)-1]
	mine := BlockChain.DigMine(currentBlock, difficulty)
	c.JSON(http.StatusOK,gin.H{
		"message":"挖到矿了",
		"block":mine.String(),
	})
}

func NewTransactionHandler(c *gin.Context) {
	var t *BlockChain.Transaction
	e:= c.ShouldBind(t)
	if e != nil {
		panic("参数错误")
	}

	transaction := &BlockChain.Transaction{
		From:   t.From,
		To:     t.To,
		Amount: t.Amount,
	}
	c.JSON(http.StatusOK, gin.H{
		"message":     "创建交易成功",
		"transaction": transaction.String(),
	})

}

func ChainHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"当前区块长度": len(blockChain.Blocks),
		"当前区块链":  blockChain,
	})
}

func NodesRegisterHandler(c *gin.Context) {
	nodes := c.Param("nodes")
	if nodes == "" || len(nodes) == 0{
		panic("参数错误")
	}
	newNodes := strings.Split(nodes,",")
	for _,node :=range newNodes{
		blockChain.RegisterNode(node)
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"新节点添加完成",
		"总结点": ToString(blockChain.Nodes),
	})
}

func NodesResolveHandler(c *gin.Context) {
	resolved := BlockChain.ResolveConflicts(blockChain)
	if resolved {
		c.JSON(http.StatusOK, gin.H{
			"message": "区块链已经解决冲突",
			"当前的区块链":  blockChain.String(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "区块链是最新的",
			"当前的区块链":  blockChain.String(),
		})
	}
}
