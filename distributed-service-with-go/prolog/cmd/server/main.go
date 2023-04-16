package main

import (
	"log"

	"github.com/iamminji/go-examples/distributed-service-with-go/prolog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
