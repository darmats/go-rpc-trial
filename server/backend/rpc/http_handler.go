package rpc

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/darmats/go-rpc-trial/server/backend/rpc/rpchttp"
)

type HTTPHandler struct {
	Logger *log.Logger
	Server *http.Server
}

func (h *HTTPHandler) ListenAndServe(address string) error {
	h.Logger = log.New(os.Stdout, "[HTTP] ", log.Ldate|log.Lmicroseconds)

	if len(address) == 0 {
		return errors.New("empty address")
	}

	hello := &rpchttp.Hello{h.Logger}
	http.HandleFunc("/hello", hello.Hello)

	h.Logger.Printf("HTTP listen start on %v\n", address)

	h.Server = &http.Server{Addr: address}
	return h.Server.ListenAndServe()
}

func (h *HTTPHandler) Shutdown(ctx context.Context) error {
	if ctx == nil {
		ctx = context.Background()
	}
	if h.Server == nil {
		return errors.New("not running")
	}

	h.Logger.Println("HTTP shutdown...")
	return h.Server.Shutdown(ctx)
}
