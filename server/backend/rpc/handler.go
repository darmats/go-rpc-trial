package rpc

import "context"

type Handler interface {
	ListenAndServe(address string) error
	Shutdown(ctx context.Context) error
}
