package p2p

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/acetimesolutions/blockchain-golang/blockchain"
	"github.com/acetimesolutions/blockchain-golang/config"
	"github.com/gin-gonic/gin"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type P2pServer struct {
	blockchain blockchain.Blockchain
	sockets    []string
}

var conf config.Config

func init() {
	conf.LoadConfigs()

	websocketConnectPeers()
	websocketHandler(nil, &http.Request{})
}

func (p *P2pServer) Run(e *gin.Engine, bc blockchain.Blockchain) {
	e.LoadHTMLFiles("static/index.html")

	e.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	e.GET("/p2p/connect", func(c *gin.Context) {
		websocketHandler(c.Writer, c.Request)
	})

	// e.GET("/p2p/connect-peers", func(c *gin.Context) {
	// 	websocketConnectPeers()
	// })
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {

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

	err = wsjson.Write(ctx, conn, "Socket connected")
	if err != nil {
		fmt.Print(err)
	}

	var v interface{}
	err = wsjson.Read(ctx, conn, &v)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(v)

	conn.Close(websocket.StatusNormalClosure, "")
}

func websocketConnectPeers() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	peers := conf.Peers

	for i := 0; i < len(peers); i++ {
		peer := peers[i]

		c, _, err := websocket.Dial(ctx, peer, nil)
		if err != nil {
			fmt.Print("Fail in connect to peer")
			log.Fatal(err)
		}
		defer c.Close(websocket.StatusInternalError, "the sky is falling")

		err = wsjson.Write(ctx, c, "Socket connected")
		if err != nil {
			fmt.Print(err)
			log.Fatal(err)
		}

		var v interface{}
		err = wsjson.Read(ctx, c, &v)
		if err != nil {
			fmt.Print(err)
			log.Fatal(err)
		}
		fmt.Print(v)

		c.Close(websocket.StatusNormalClosure, "")
	}
}
