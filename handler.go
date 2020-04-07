package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func MineHandler(c *gin.Context) {

	r := blockChain.Validate()
	if !r {
		panic("该区块链校验失败")
	}
	preHash:= blockChain.Blocks[len(blockChain.Blocks)-1].Hash
	newBlock := NewBlock(transactionPool,preHash)
	mine := DigMine(newBlock, difficulty)
	blockChain.Blocks = append(blockChain.Blocks, newBlock)
	blockChain.TransactionPool=nil
	c.JSON(http.StatusOK, gin.H{
		"message": "挖到矿了",
		"block":   mine,
	})
}

func NewTransactionHandler(c *gin.Context) {
	var t Transaction
	e := c.ShouldBindJSON(&t)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "参数错误:" + e.Error(),
		})
		return
	}

	transaction := &Transaction{
		From:   t.From,
		To:     t.To,
		Amount: t.Amount,
	}
	//blockChain.TransactionPool = append(blockChain.TransactionPool, transaction)
	transactionPool = append(blockChain.TransactionPool, transaction)
	blockChain.TransactionPool = transactionPool
	c.JSON(http.StatusOK, gin.H{
		"message":     "创建交易成功",
		"transaction": transaction,
	})

}



func ChainHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"length": len(blockChain.Blocks),
		"chain":  blockChain,
	})
}

type Neighbor struct {
	Nodes string `json:"nodes"`
}

func NodesRegisterHandler(c *gin.Context) {

	var neighbor Neighbor
	e := c.ShouldBindJSON(&neighbor)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "参数错误:" + e.Error(),
		})
		return
	}

	newNodes := strings.Split(neighbor.Nodes, ",")
	for _, node := range newNodes {
		blockChain.RegisterNode(node)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "新节点添加完成",
		"nodes":     ToString(blockChain.Nodes),
	})
}

func NodesResolveHandler(c *gin.Context) {
	resolved := ResolveConflicts(blockChain)
	if resolved {
		c.JSON(http.StatusOK, gin.H{
			"message": "区块链已经解决冲突",
			"chain":  blockChain,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "区块链是最新的",
			"chain":  blockChain,
		})
	}
}
