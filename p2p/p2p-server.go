package p2p

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type P2pServer struct {
}

func (p *P2pServer) Run(rg *gin.RouterGroup) {
	rg.GET("/websocket", p.websocketHandler())
}

func (p *P2pServer) websocketHandler() gin.HandlerFunc {
	// websocket.Accept()

	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, []string{})
	}
}
