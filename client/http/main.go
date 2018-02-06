package http

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/darmats/go-rpc-trial/define"
)

func Run1(loop, wait int) error {
	url := fmt.Sprintf(`%s/hello?name=%s&wait=%d`, define.BackendHTTPEndPoint, "World", wait)
	client := http.Client{}

	for i := 0; i < loop; i++ {
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return err
		}
		res, err := client.Do(req)
		if err != nil {
			return err
		}
		res.Body.Close()
	}

	return nil
}

func Run2(loop, wait int) error {
	url := fmt.Sprintf(`%s/hello?name=%s&wait=%d`, define.BackendHTTPEndPoint, "World", wait)
	client := http.Client{}

	e := make(chan error)
	wg := &sync.WaitGroup{}
	for i := 0; i < loop; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			req, err := http.NewRequest(http.MethodGet, url, nil)
			if err != nil {
				e <- err
				return
			}
			res, err := client.Do(req)
			if err != nil {
				e <- err
				return
			}
			res.Body.Close()
		}()
	}

	go func() {
		wg.Wait()
		e <- nil
	}()

	return <-e
}
