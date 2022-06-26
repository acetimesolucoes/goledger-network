package p2p

import (
	"bytes"
	"context"
	"encoding/json"
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
	Blockchain  blockchain.Blockchain
	Connection  *websocket.Conn
	Connections []*websocket.Conn
	Contexts    []*context.Context
	Config      config.Config
}

func (p *P2pServer) Run(e *gin.Engine, bc blockchain.Blockchain) {

	p.Config.LoadConfigs()

	p.Blockchain.ReplaceChain(bc.Chain)
	p.websocketHandler(nil, &http.Request{})
	p.connectToPeers()

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
	p.Contexts = append(p.Contexts, &ctx)
	p.Connections = append(p.Connections, conn)

	defer cancel()

	p.messageHandler(&ctx, conn)
	p.SyncChains()

	conn.Close(websocket.StatusNormalClosure, "")
}

func (p *P2pServer) connectToPeers() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	for i := 0; i < len(p.Config.Peers); i++ {
		peer := p.Config.Peers[i]

		conn, _, err := websocket.Dial(ctx, peer, nil)
		if err != nil {
			fmt.Print("Fail in connect to peer\n")
			log.Fatal(err)
		}
		defer conn.Close(websocket.StatusInternalError, "the sky is falling")

		// p.Connections = append(p.Connections, *conn)

		err = wsjson.Write(ctx, *&conn, "Slave -> Master \n")
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

func (p *P2pServer) messageHandler(cxt *context.Context, conn *websocket.Conn) {

	str := ObjectToString(p.Blockchain.Chain)

	err := wsjson.Write(*cxt, conn, str)
	if err != nil {
		log.Fatal(err)
	}

	// conn.Close(websocket.StatusNormalClosure, "")
	fmt.Print(str)

	if str == "null" {
		return
	}

	// obj := StringToObject[[]blockchain.Block](str)
	// fmt.Print(obj)
}

func ObjectToString[T any](object T) string {
	byteArray, err := json.Marshal(object)

	if err != nil {
		fmt.Print(err)
	}

	str := string([]byte(byteArray))

	return str
}

func StringToObject[T any](str string) T {

	bytes := bytes.NewBufferString(str)
	var object T

	err := json.Unmarshal(bytes.Bytes(), &object)
	if err != nil {
		log.Fatal(err)
	}

	return object
}

func (p *P2pServer) sendChain(ctx *context.Context, conn websocket.Conn) {
	str := ObjectToString(p.Blockchain.Chain)

	if str == "null" {
		return
	}

	fmt.Print(str)

	err := wsjson.Write(*ctx, &conn, str)
	if err != nil {
		log.Fatal(err)
	}
}

func (p *P2pServer) SyncChains() {
	for i, conn := range p.Connections {
		p.sendChain(p.Contexts[i], *conn)
	}
}
