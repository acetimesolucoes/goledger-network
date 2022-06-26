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

func (b *BlockchainServer) Run(rg *gin.RouterGroup) {

	rg.GET("/blocks", b.handleBlocks())
	rg.POST("/blocks", b.handleMine())
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
