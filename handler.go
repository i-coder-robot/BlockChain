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
		"block":   mine.String(),
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
	if nodes == "" || len(nodes) == 0 {
		panic("参数错误")
	}
	newNodes := strings.Split(nodes, ",")
	for _, node := range newNodes {
		blockChain.RegisterNode(node)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "新节点添加完成",
		"总结点":     ToString(blockChain.Nodes),
	})
}

func NodesResolveHandler(c *gin.Context) {
	resolved := ResolveConflicts(blockChain)
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
