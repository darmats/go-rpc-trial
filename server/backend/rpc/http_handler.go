package rpc

import (
	"context"
	"log"
	"net/http"

	"github.com/darmats/go-rpc-trial/server/backend/rpc/rpchttp"
)

type HTTPHandler struct {
	Server *http.Server
}

func (h *HTTPHandler) ListenAndServe(address string) error {
	if len(address) == 0 {
		address = ":8080"
	}

	http.HandleFunc("/hello", rpchttp.Hello)

	log.Printf("HTTP listen start on %v\n", address)

	h.Server = &http.Server{Addr: address}
	return h.Server.ListenAndServe()
}

func (h *HTTPHandler) Shutdown(ctx context.Context) error {
	if ctx == nil {
		ctx = context.Background()
	}

	//log.Println("HTTP shutdown...")
	return h.Server.Shutdown(ctx)
}
