package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/acetimesolutions/blockchain-golang/blockchain"
	"github.com/acetimesolutions/blockchain-golang/p2p"
	"github.com/gin-gonic/gin"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = ":3001"
	}

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		new(blockchain.BlockchainServer).Run(v1)
		new(p2p.P2pServer).Run(v1)
	}

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"version":      "1.0.0",
			"package_name": "@acetime/blockchain",
			"timestamp":    time.Now().UnixNano(),
		})
	})

	router.Run(":" + PORT)
}
