package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/acetimesolutions/blockchain-golang/blockchain"
	"github.com/acetimesolutions/blockchain-golang/config"
	"github.com/acetimesolutions/blockchain-golang/p2p"
	"github.com/gin-gonic/gin"
)

var conf config.Config
var bc blockchain.Blockchain

func main() {
	router := gin.Default()

	var p2pServer p2p.P2pServer
	p2pServer.Run(router, bc)

	var blockchainServer blockchain.BlockchainServer
	blockchainServer.Run(router)

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"version":      "1.0.0",
			"package_name": "@acetime/blockchain",
			"timestamp":    time.Now().UnixNano(),
		})
	})

	router.Run(":" + strconv.Itoa(conf.HTTP_PORT))
}

func init() {
	conf.LoadConfigs()
	bc.Init()
}
