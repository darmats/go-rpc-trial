package rpchttp

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Hello(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

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
	writer.WriteHeader(http.StatusOK)
	writer.Write(b)
}
