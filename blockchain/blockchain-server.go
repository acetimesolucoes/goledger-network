package blockchain

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlockchainServer struct {
	Blockchain Blockchain
	// P2PServer  *p2p.P2pServer
}

func (b *BlockchainServer) Run(e *gin.Engine, bc Blockchain) {

	b.Blockchain.ReplaceChain(bc.Chain)

	e.GET("/block", b.handleBlocks())
	e.POST("/mine", b.handleMine())
}

func (b *BlockchainServer) handleBlocks() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, b.Blockchain.Chain)
	}
}

func (b *BlockchainServer) handleMine() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var block any
		ctx.BindJSON(&block)
		b.Blockchain.AddBlock(block)

		// b.P2PServer.SyncChains()

		ctx.JSON(http.StatusOK, b.Blockchain.Chain)
	}
}
