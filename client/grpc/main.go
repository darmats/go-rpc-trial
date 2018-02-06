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

func main() {
	os.Exit(run())
}

func run() int {

	var mode, loop, wait int

	flag.IntVar(&mode, "m", 1, "")
	flag.IntVar(&loop, "l", 10000, "loop count")
	flag.IntVar(&wait, "w", 0, "wait (millisecond)")
	flag.Parse()

	var err error

	log.Printf("method: Run%d(), loop: %d, wait: %d millisecond", mode, loop, wait)

	start := time.Now()

	switch mode {
	case 1:
		err = Run1(loop, wait)
	case 2:
		err = Run2(loop, wait)
	case 3:
		err = Run3(loop, wait)
	case 4:
		err = Run4(loop, wait)
	}

	d := time.Since(start)

	if err != nil {
		log.Println(err)
		return 1
	}

	log.Printf("duration: %v", d)

	return 0
}

func Run1(loop, wait int) error {
	conn, err := grpc.Dial(":"+define.BackendGRPCPort, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	client := pb.NewHelloClient(conn)

	for i := 0; i < loop; i++ {
		_, err = client.Say(context.Background(), &pb.HelloRequest{Name: "World", Wait: int32(wait)})
		if err != nil {
			return err
		}
	}

	return nil
}

func Run2(loop, wait int) error {
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

			_, err = client.Say(context.Background(), &pb.HelloRequest{Name: "World", Wait: int32(wait)})
			if err != nil {
				//e <- err
				return
			}
		}()
	}
	wg.Wait()

	return nil
}

func Run3(loop, wait int) error {
	for i := 0; i < loop; i++ {
		conn, err := grpc.Dial(":"+define.BackendGRPCPort, grpc.WithInsecure())
		if err != nil {
			return err
		}

		client := pb.NewHelloClient(conn)
		_, err = client.Say(context.Background(), &pb.HelloRequest{Name: "World", Wait: int32(wait)})
		if err != nil {
			conn.Close()
			return err
		}
		conn.Close()
	}

	return nil
}

func Run4(loop, wait int) error {
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
			_, err = client.Say(context.Background(), &pb.HelloRequest{Name: "World", Wait: int32(wait)})
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
