package main

import (
	"context"
	"flag"
	"log"
	"os"
	"sync"
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
	flag.IntVar(&loop, "l", 10000, "")
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
		_, err = client.Say(context.Background(), &pb.HelloRequest{Name: "World"})
		if err != nil {
			return err
		}
	}

	return nil
}

func run2() error {
	conn, err := grpc.Dial(":"+define.BackendGRPCPort, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	client := pb.NewHelloClient(conn)

	// todo: receive err
	//e := make(chan error)
	wg := &sync.WaitGroup{}
	for i := 0; i < loop; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			_, err = client.Say(context.Background(), &pb.HelloRequest{Name: "World"})
			if err != nil {
				//e <- err
				return
			}
		}()
	}
	wg.Wait()

	return nil
}

func run3() error {
	for i := 0; i < loop; i++ {
		conn, err := grpc.Dial(":"+define.BackendGRPCPort, grpc.WithInsecure())
		if err != nil {
			return err
		}

		client := pb.NewHelloClient(conn)
		_, err = client.Say(context.Background(), &pb.HelloRequest{Name: "World"})
		if err != nil {
			conn.Close()
			return err
		}
		conn.Close()
	}

	return nil
}

func run4() error {
	// todo: receive err
	wg := &sync.WaitGroup{}
	for i := 0; i < loop; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			conn, err := grpc.Dial(":"+define.BackendGRPCPort, grpc.WithInsecure())
			if err != nil {
				return
			}

			client := pb.NewHelloClient(conn)
			_, err = client.Say(context.Background(), &pb.HelloRequest{Name: "World"})
			if err != nil {
				conn.Close()
				return
			}
			conn.Close()
		}()
	}
	wg.Wait()

	return nil
}
