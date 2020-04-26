package main

import (
	"github.com/FanyCastro/go-myfirstrestservice/internal/app/server"
	"log"
)

func run (port int) error {
	srv := server.NewServer(port)
	return srv.Start()
}

func main() {
	err := run(8080)
	if err != nil {
		log.Fatalf("Fail to start server, %q", err.Error())
	}
}
