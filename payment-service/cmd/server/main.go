package main

import (
	"log"
	"net/http"

	"github.com/imrenagi/microservice-demo/payment-service/web"
)

func main() {
	s := web.NewServer()
	if err := http.ListenAndServe(":8081", s.Router); err != nil {
		log.Fatalf("Server can't run. Got: `%v`", err)
	}
}
