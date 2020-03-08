package main

import (
	"flag"
	"github.com/gin-gonic/gin"
)

func main() {
	r:= gin.Default()
	r.GET("/mine",MineHandler)
	r.POST("/transactions/new",NewTransactionHandler)
	r.GET("/chain",ChainHandler)
	r.POST("/nodes/register",NodesRegisterHandler)
	r.POST("/nodes/resolve",NodesResolveHandler)

	port:=flag.String("port","8081","请输入端口号")
	r.Run("0.0.0.0:"+*port)
}

