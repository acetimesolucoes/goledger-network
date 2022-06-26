package main

import (
	"net/http"
	"time"

	"github.com/acetimesolutions/blockchain-golang/blockchain"
	"github.com/gin-gonic/gin"
)

func main() {

	var bc blockchain.Blockchain
	bc.Init()

	route := gin.Default()

	route.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message":   "pong",
			"datetime":  time.Now().UTC(),
			"timestamp": time.Now().UnixNano(),
		})
	})

	route.GET("/blocks", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, bc.Chain)
	})

	route.POST("/mine", func(ctx *gin.Context) {

		var block any
		ctx.BindJSON(&block)

		bc.AddBlock(block)

		ctx.JSON(http.StatusOK, bc.Chain)
	})

	route.Run(":3001")
}
