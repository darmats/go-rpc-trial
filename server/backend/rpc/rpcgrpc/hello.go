package rpcgrpc

import (
	"fmt"
	"log"
	"time"

	"github.com/darmats/go-rpc-trial/define/grpc/pb"
	"golang.org/x/net/context"
)

type Hello struct {
	Logger *log.Logger
}

func (h *Hello) Say(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	if req.Wait == 1 {
		time.Sleep(500 * time.Millisecond)
	}

	res := &pb.HelloResponse{}
	name := req.Name
	if len(name) == 0 {
		h.Logger.Println(`"name" is empty`)
		name = "user"
	}
	res.Message = fmt.Sprintf("Hello, %s!", name)

	return res, nil
}
