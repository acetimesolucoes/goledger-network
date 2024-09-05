package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	HTTP_PORT int
	P2P_PORT  int
	Peers     []string
}

func init() {
	err := godotenv.Load("config/env/.env")

	if err != nil {
		log.Fatal(err)
	}
}

func (c *Config) LoadConfigs() {
	c.getP2PPort()
	c.getPeers()
	c.getPort()
}

func (c *Config) getPort() {
	port, err := strconv.Atoi(strings.TrimSpace(os.Getenv("HTTP_PORT")))

	if err != nil {
		port = 3003
	}

	c.HTTP_PORT = port
}

func (c *Config) getPeers() {

	peers := strings.TrimSpace(os.Getenv("PEERS"))

	if peers != "" {
		println("PEERS=", peers)
		c.Peers = strings.Split(peers, ",")
	} else {
		fmt.Println("PEERS=", peers)
		c.Peers = []string{"ws://localhost:3001/p2p/connect", "ws://localhost:3002/p2p/connect"}
		// c.Peers = []string{}
	}

}

func (c *Config) getP2PPort() {
	p2pPort, err := strconv.Atoi(strings.TrimSpace(os.Getenv("P2P_PORT")))

	if err != nil {
		p2pPort = 5003
	}

	c.P2P_PORT = p2pPort
}

func (c *Config) ToString() {
	fmt.Printf(`
		Config file -
		HTTP_PORT: %d
		P2P_PORT %d
		PEERS %q
	`, c.HTTP_PORT, c.P2P_PORT, c.Peers)
}
