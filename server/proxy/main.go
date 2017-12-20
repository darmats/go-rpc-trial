package main

import (
	"os"

	"github.com/darmats/go-rpc-trial/server/proxy/router"
	"github.com/gin-gonic/gin"
)

func main() {
	os.Exit(run())
}

func run() int {

	g := gin.Default()
	route(g)

	g.Run(":8088")

	return 0
}

func route(g *gin.Engine) {
	hello := &router.Hello{}
	g.GET("/hello/http", hello.HTTP)
	g.GET("/hello/grpc", hello.GRPC)
}
