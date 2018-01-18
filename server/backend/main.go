package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/darmats/go-rpc-trial/define"
	"github.com/darmats/go-rpc-trial/server/backend/rpc"
)

var rpcHandlers = []rpc.Handler{
	&rpc.HTTPHandler{},
	&rpc.GRPCHandler{},
}

func main() {
	os.Exit(run())
}

func run() int {

	sig := make(chan os.Signal, 1)

	for _, rpcHandler := range rpcHandlers {
		go func(handler rpc.Handler) {
			if err := handler.ListenAndServe(address(handler)); err != nil {
				log.Println(err)
			}
		}(rpcHandler)
	}

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	for _, rpcHandler := range rpcHandlers {
		rpcHandler.Shutdown(ctx)
	}

	return 0
}

func address(handler rpc.Handler) string {
	switch handler.(type) {
	case *rpc.HTTPHandler:
		return ":8080"
	case *rpc.GRPCHandler:
		return ":" + define.BackendGRPCPort
	default:
		return ""
	}
}
