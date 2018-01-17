package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"os"
	"time"

	"github.com/darmats/go-rpc-trial/define"
	"github.com/darmats/go-rpc-trial/define/grpc/pb"
	"google.golang.org/grpc"
)

var (
	mode int
	loop int
)

func main() {
	os.Exit(run())
}

func run() int {

	flag.IntVar(&mode, "m", 1, "")
	flag.IntVar(&loop, "l", 100000, "")
	flag.Parse()

	var err error

	start := time.Now()

	switch mode {
	case 1:
		err = run1()
	case 2:
		err = run2()
	case 3:
		err = run3()
	case 4:
		err = run4()
	}

	d := time.Since(start)

	if err != nil {
		log.Println(err)
		return 1
	}

	log.Println(d)

	return 0
}

func run1() error {
	conn, err := grpc.Dial(":"+define.BackendGRPCPort, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	client := pb.NewHelloClient(conn)

	for i := 0; i < loop; i++ {
		_, err := client.Say(context.Background(), &pb.HelloRequest{Name: "World"})
		if err != nil {
			return err
		}
	}

	return nil
}

func run2() error {
	return errors.New("not implemented")
}

func run3() error {
	return errors.New("not implemented")
}

func run4() error {
	return errors.New("not implemented")
}
