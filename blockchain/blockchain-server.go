package blockchain

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlockchainServer struct {
}

var bc Blockchain

func init() {
	bc.Init()
}

func (b *BlockchainServer) Run(e *gin.Engine) {
	e.GET("/block", b.handleBlocks())
	e.POST("/mine", b.handleMine())
}

func (b *BlockchainServer) handleBlocks() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, bc.Chain)
	}
}

func (b *BlockchainServer) handleMine() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var block any
		ctx.BindJSON(&block)
		bc.AddBlock(block)
		ctx.JSON(http.StatusOK, bc.Chain)
	}
}
