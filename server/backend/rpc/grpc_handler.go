package rpc

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/darmats/go-rpc-trial/server/backend/rpc/rpcgrpc"
	"github.com/darmats/go-rpc-trial/server/backend/rpc/rpcgrpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCHandler struct {
	Server *grpc.Server
}

func (h *GRPCHandler) ListenAndServe(address string) error {
	if len(address) == 0 {
		address = ":50051"
	}

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to listen: %v", err))
	}

	srv := grpc.NewServer()

	pb.RegisterHelloServer(srv, &rpcgrpc.Hello{})
	reflection.Register(srv)

	log.Printf("gRPC listen start on %v\n", address)

	h.Server = srv
	if err := srv.Serve(listener); err != nil {
		return errors.New(fmt.Sprintf("failed to serve: %v", err))
	}

	return nil
}

func (h *GRPCHandler) Shutdown(ctx context.Context) error {
	//log.Println("gRPC shutdown...")
	h.Server.GracefulStop()
	return nil
}
