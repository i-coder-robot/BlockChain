package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"time"
)

const (
	rewards    = 50.0
	difficulty = 4
)

var blockChain *BlockChain

//var transactions []*Transaction
var transactionPool []*Transaction

func Init() {
	transaction := &Transaction{
		From:   "",
		To:     "欢喜哥",
		Amount: rewards,
	}

	transactions := []*Transaction{transaction}
	transactionPool = []*Transaction{}
	block := &Block{
		Transactions: transactions,
		PreBlockHash: "",
		Hash:         "",
		Nonce:        0,
		TimeStamp:    time.Now().Unix(),
	}
	block.Hash = ProofOfWorkWithDifficult(block, difficulty)
	blockChain = NewBlockChain(transactions, block, difficulty, rewards)
}

func main() {
	r := gin.Default()
	r.GET("/mine", MineHandler)
	r.POST("/transactions/new", NewTransactionHandler)
	r.GET("/chain", ChainHandler)
	r.POST("/nodes/register/", NodesRegisterHandler)
	r.POST("/nodes/resolve", NodesResolveHandler)
	Init()
	port := flag.String("port", "", "请输入端口号")
	flag.Parse()
	portStr := *port
	if portStr == "" {
		portStr = "8081"
	}
	r.Run("0.0.0.0:" + portStr)
}
