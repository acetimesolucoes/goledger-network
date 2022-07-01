package p2p

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/acetimesolutions/goledger-network/blockchain"
	"github.com/acetimesolutions/goledger-network/config"
	"github.com/gin-gonic/gin"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type IP2pServer interface {
	Run()
	SyncChains()
	P2pServer
}

type P2pServer struct {
	Blockchain  *blockchain.Blockchain
	Connections []websocket.Conn
	Contexts    []context.Context
	Config      config.Config
	Context     *context.Context
	Connection  *websocket.Conn
}

func (p *P2pServer) Run(e *gin.Engine, b *blockchain.Blockchain) {

	p.Blockchain = b

	e.LoadHTMLFiles("static/index.html")

	e.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	e.GET("/p2p/connect", func(c *gin.Context) {
		p.websocketHandler(c.Writer, c.Request)
	})

	// p.websocketHandler(nil, &http.Request{})
	p.connectToPeers()
}

func (p *P2pServer) websocketHandler(w http.ResponseWriter, r *http.Request) {

	if w == nil || r == nil {
		return
	}

	conn, err := websocket.Accept(w, r, nil)

	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v", err)
		return
	} else {
		fmt.Println("start websocket connection...")
	}

	ctx, _ := context.WithTimeout(r.Context(), time.Second*100)
	// defer cancel()

	p.Contexts = append(p.Contexts, ctx)
	p.Connections = append(p.Connections, *conn)

	p.messageHandler(&p.Contexts[0], &p.Connections[0])

	// defer conn.Close(websocket.StatusInternalError, "closed websocket connection...")
	// defer cancel()
	// conn.Close(websocket.StatusNormalClosure, "")
}

func (p *P2pServer) connectToPeers() {

	ctx, _ := context.WithTimeout(context.Background(), time.Minute)
	// defer cancel()

	for i := 0; i < len(p.Config.Peers); i++ {
		peer := p.Config.Peers[i]

		conn, _, err := websocket.Dial(ctx, peer, nil)
		if err != nil {
			fmt.Println("Fail in connect to peer")
			log.Panic(err)
		}
		// defer conn.Close(websocket.StatusInternalError, "the sky is falling")

		p.Contexts = append(p.Contexts, ctx)
		p.Connections = append(p.Connections, *conn)

		err = wsjson.Write(ctx, conn, "New peer connected to server \n")
		if err != nil {
			fmt.Println(err)
			// log.Fatal(err)
		}

		var jsonReceived string
		err = wsjson.Read(ctx, conn, &jsonReceived)
		if err != nil {
			fmt.Println(err)
			// log.Fatal(err)
		}

		fmt.Println("received: ", jsonReceived)

		fmt.Println(StringToObject[[]blockchain.Block](jsonReceived))

		chainToReplace := StringToObject[[]blockchain.Block](jsonReceived)

		p.Blockchain.ReplaceChain(chainToReplace)

		// conn.Close(websocket.StatusNormalClosure, "")
	}
}

func (p *P2pServer) messageHandler(ctx *context.Context, conn *websocket.Conn) {

	str := ObjectToString(p.Blockchain.Chain)

	err := wsjson.Write(*ctx, conn, str)
	if err != nil {
		log.Fatal(err)
	}

	if str == "null" {
		return
	}

	var jsonReceived string
	err = wsjson.Read(*ctx, conn, &jsonReceived)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("received: ", jsonReceived)
}

func ObjectToString[T any](object T) string {
	byteArray, err := json.Marshal(object)

	if err != nil {
		fmt.Println(err)
	}

	str := string([]byte(byteArray))

	return str
}

func StringToObject[T any](str string) T {

	bytes := bytes.NewBufferString(str)
	var object T

	err := json.Unmarshal(bytes.Bytes(), &object)
	if err != nil {
		// log.Fatal(err)
	}

	return object
}

func (p *P2pServer) sendChain(ctx context.Context, conn websocket.Conn) {
	// ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	// defer cancel()

	str := ObjectToString(p.Blockchain.Chain)

	if str == "null" {
		return
	}

	fmt.Println("json to send: " + str)

	err := wsjson.Write(ctx, &conn, str)
	// err := wsjson.Write(ctx, conn, p.Blockchain.Chain)
	if err != nil {
		log.Fatal(err)
	}

	var v interface{}
	err = wsjson.Read(ctx, &conn, &v)
	if err != nil {
		log.Fatal(err)
	}
}

func (p *P2pServer) SyncChains() {
	for i, _ := range p.Connections {
		p.sendChain(p.Contexts[i], p.Connections[i])
	}
}
