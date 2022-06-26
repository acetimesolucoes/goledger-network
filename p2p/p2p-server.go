package p2p

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/acetimesolutions/blockchain-golang/blockchain"
	"github.com/acetimesolutions/blockchain-golang/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/internal/bytesconv"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type P2pServer struct {
	blockchain blockchain.Blockchain
	// sockets    []string
}

var conf config.Config
var websocketConnection *websocket.Conn

func init() {
	var instance P2pServer

	conf.LoadConfigs()

	instance.connectToPeers()
	instance.websocketHandler(nil, &http.Request{})
}

func (p *P2pServer) Run(e *gin.Engine, bc blockchain.Blockchain) {

	p.blockchain = bc

	e.LoadHTMLFiles("static/index.html")

	e.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	e.GET("/p2p/connect", func(c *gin.Context) {
		p.websocketHandler(c.Writer, c.Request)
	})
}

func (p *P2pServer) websocketHandler(w http.ResponseWriter, r *http.Request) {

	if w == nil || r == nil {
		return
	}

	conn, err := websocket.Accept(w, r, nil)

	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v", err)
		return
	}
	defer conn.Close(websocket.StatusInternalError, "closed websocket connection...")

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	err = wsjson.Write(ctx, conn, "Master -> Slave \n")
	if err != nil {
		fmt.Print(err)
	}

	p.handleMessage(p.blockchain.Chain)

	var v interface{}
	err = wsjson.Read(ctx, conn, &v)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(v)

	conn.Close(websocket.StatusNormalClosure, "")
}

func (p *P2pServer) connectToPeers() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	peers := conf.Peers

	for i := 0; i < len(peers); i++ {
		peer := peers[i]

		conn, _, err := websocket.Dial(ctx, peer, nil)
		if err != nil {
			fmt.Print("Fail in connect to peer\n")
			log.Fatal(err)
		}
		defer conn.Close(websocket.StatusInternalError, "the sky is falling")

		websocketConnection = conn

		// p.handleMessage("Master -> Slave Some")
		p.handleMessage(p.blockchain.Chain)

		err = wsjson.Write(ctx, conn, "Slave -> Master \n")
		if err != nil {
			fmt.Print(err)
			log.Fatal(err)
		}

		var v interface{}
		err = wsjson.Read(ctx, conn, &v)
		if err != nil {
			fmt.Print(err)
			log.Fatal(err)
		}
		fmt.Print(v)

		conn.Close(websocket.StatusNormalClosure, "")
	}
}

func (p *P2pServer) handleMessage(message interface{}) {

	// data := gin.H{"data": message}
	bytes, _ := json.Marshal(message)

	str := bytesconv.BytesToString(bytes)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	err := wsjson.Write(ctx, websocketConnection, str)

	if err != nil {
		fmt.Print(err)
		log.Fatal(err)
	}
}
