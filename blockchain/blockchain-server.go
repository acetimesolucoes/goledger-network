package blockchain

import (
	"net/http"
	"strconv"
	"time"

	"github.com/acetimesolutions/goledger-network/config"
	"github.com/gin-gonic/gin"
)

type BlockchainServer struct {
	Blockchain *Blockchain
	Config     config.Config
}

func (b *BlockchainServer) Run(e *gin.Engine, bc *Blockchain) {

	b.Blockchain = bc

	e.GET("/blocks", b.handleBlocks())
	e.POST("/mine", b.handleMine())

	e.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"version":      "1.0.0",
			"package_name": "@acetime/blockchain",
			"timestamp":    time.Now().UnixNano(),
		})
	})

	e.Run(":" + strconv.Itoa(b.Config.HTTP_PORT))
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
		b.Blockchain.AddBlock(&block)

		// ctx.Redirect(http.StatusPermanentRedirect, "p2p/sync")
		ctx.JSON(http.StatusOK, b.Blockchain.Chain)
	}
}
