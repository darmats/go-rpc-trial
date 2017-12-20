package rpcgrpc

import (
	"fmt"
	"log"

	"github.com/darmats/go-rpc-trial/server/backend/rpc/rpcgrpc/pb"
	"golang.org/x/net/context"
)

type Hello struct {
}

func (h *Hello) Say(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	res := &pb.HelloResponse{}
	name := req.Name
	if len(name) == 0 {
		log.Println(`"name" is empty`)
		name = "user"
	}
	res.Message = fmt.Sprintf("Hello, %s!", name)

	return res, nil
}
