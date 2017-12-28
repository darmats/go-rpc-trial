package rpchttp

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Hello struct {
	Logger *log.Logger
}

func (h *Hello) Hello(writer http.ResponseWriter, req *http.Request) {
	defer func(start time.Time) {
		h.Logger.Printf("| %v | %s | %s | %s\n",
			time.Since(start),
			req.RemoteAddr,
			req.Method,
			req.RequestURI,
			//req.URL.Path,
		)
	}(time.Now())

	wait := req.URL.Query().Get("wait")
	if wait == "1" {
		time.Sleep(time.Second)
	}

	res := struct {
		Message string `json:"message"`
	}{}

	name := req.URL.Query().Get("name")
	if len(name) == 0 {
		log.Println(`"name" is empty`)
		name = "user"
	}
	res.Message = fmt.Sprintf("Hello, %s!", name)
	b, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(b)
}
