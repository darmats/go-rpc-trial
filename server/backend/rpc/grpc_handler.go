package rpc

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/darmats/go-rpc-trial/define/grpc/pb"
	"github.com/darmats/go-rpc-trial/server/backend/rpc/rpcgrpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type GRPCHandler struct {
	Logger *log.Logger
	Server *grpc.Server
}

func (h *GRPCHandler) ListenAndServe(address string) error {
	h.Logger = log.New(os.Stdout, "[gRPC] ", log.Ldate|log.Lmicroseconds)

	if len(address) == 0 {
		return errors.New("empty address")
	}

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to listen: %v", err))
	}

	srv := grpc.NewServer(
		grpc.UnaryInterceptor(h.unaryInterceptor()),
	)

	pb.RegisterHelloServer(srv, &rpcgrpc.Hello{h.Logger})
	reflection.Register(srv)

	h.Logger.Printf("gRPC listen start on %v\n", address)

	h.Server = srv
	if err := srv.Serve(listener); err != nil {
		return errors.New(fmt.Sprintf("failed to serve: %v", err))
	}

	return nil
}

func (h *GRPCHandler) Shutdown(ctx context.Context) error {
	if h.Server == nil {
		return errors.New("not running")
	}

	h.Logger.Println("gRPC shutdown...")
	h.Server.GracefulStop()
	return nil
}

func (h *GRPCHandler) unaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func(start time.Time) {
			d := time.Since(start)

			if r := recover(); r != nil {
				err = status.Errorf(codes.Internal, "panic: %v", r)
			}

			var address string
			if p, ok := peer.FromContext(ctx); ok {
				address = p.Addr.String()
			}
			h.Logger.Printf("| %v | %s | %s\n",
				d,
				address,
				info.FullMethod,
			)
		}(time.Now())

		return handler(ctx, req)
	}
}
