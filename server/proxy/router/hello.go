package router

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/darmats/go-rpc-trial/server/proxy/define/grpc/pb"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Hello struct {
}

func (h *Hello) HTTP(ctx *gin.Context) {
	name := ctx.DefaultQuery("name", "user")

	n, err := strconv.Atoi(ctx.DefaultQuery("n", "1"))
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c, err := strconv.Atoi(ctx.DefaultQuery("c", "1"))
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	start := time.Now()

	client := http.Client{}
	for i := 0; i < n/c; i++ {
		wg := sync.WaitGroup{}
		for j := 0; j < c; j++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				req, err := http.NewRequest(http.MethodGet, EndPointHTTP+"/hello?name="+name, nil)
				if err != nil {
					ctx.AbortWithError(http.StatusInternalServerError, err)
					return
				}
				res, err := client.Do(req)
				if err != nil {
					ctx.AbortWithError(http.StatusInternalServerError, err)
					return
				}
				//defer res.Body.Close()
				res.Body.Close()
			}()
		}
		wg.Wait()
	}

	d := time.Since(start)
	ctx.Data(http.StatusOK, "text/html", []byte(d.String()+"\n"))
}

func (h *Hello) GRPC(ctx *gin.Context) {
	name := ctx.DefaultQuery("name", "user")

	n, err := strconv.Atoi(ctx.DefaultQuery("n", "1"))
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c, err := strconv.Atoi(ctx.DefaultQuery("c", "1"))
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	start := time.Now()

	con, err := grpc.Dial(EndPointGRPC, grpc.WithInsecure())
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer con.Close()

	client := pb.NewHelloClient(con)
	for i := 0; i < n/c; i++ {
		wg := &sync.WaitGroup{}
		for j := 0; j < c; j++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				_, err := client.Say(context.Background(), &pb.HelloRequest{Name: name})
				if err != nil {
					ctx.AbortWithError(http.StatusInternalServerError, err)
					return
				}
			}()
		}
		wg.Wait()
	}

	d := time.Since(start)
	ctx.Data(http.StatusOK, "text/html", []byte(d.String()+"\n"))
}
