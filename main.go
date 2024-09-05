package main

import (
	"github.com/acetimesolutions/goledger-network/blockchain"
	"github.com/acetimesolutions/goledger-network/p2p"
	"github.com/gin-gonic/gin"
)

func main() {
	var bc blockchain.Blockchain
	var p2pServer p2p.P2pServer
	var blockchainServer blockchain.BlockchainServer

	bc.Init()

	router := gin.Default()

	p2pServer.Config.LoadConfigs()
	p2pServer.Run(router, &bc)

	blockchainServer.Config.LoadConfigs()
	blockchainServer.Run(router, &bc)
}
