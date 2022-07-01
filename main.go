package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/acetimesolutions/goledger-network/blockchain"
	"github.com/acetimesolutions/goledger-network/config"
	"github.com/acetimesolutions/goledger-network/p2p"
	"github.com/gin-gonic/gin"
)

func main() {
	var conf config.Config
	var bc blockchain.Blockchain
	var p2pServer p2p.P2pServer
	var blockchainServer blockchain.BlockchainServer

	conf.LoadConfigs()
	bc.Init()

	router := gin.Default()

	p2pServer.Config.LoadConfigs()
	p2pServer.Run(router, &bc)

	blockchainServer.Run(router, &bc)

	router.GET("/blocks", handleBlocks(&blockchainServer, &p2pServer))
	router.POST("/mine", handleMine(&blockchainServer, &p2pServer))

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"version":      "1.0.0",
			"package_name": "@acetime/blockchain",
			"timestamp":    time.Now().UnixNano(),
		})
	})

	router.Run(":" + strconv.Itoa(conf.HTTP_PORT))
}

func handleBlocks(b *blockchain.BlockchainServer, p *p2p.P2pServer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, b.Blockchain.Chain)
	}
}

func handleMine(b *blockchain.BlockchainServer, p *p2p.P2pServer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var block any
		ctx.BindJSON(&block)
		b.Blockchain.AddBlock(&block)

		p.SyncChains()

		// ctx.Redirect(http.StatusPermanentRedirect, "p2p/sync")
		ctx.JSON(http.StatusOK, b.Blockchain.Chain)
	}
}
