package main

import (
	"github.com/FanyCastro/go-myfirstrestservice/internal/app/server"
	"log"
)

func main() {
	srv := server.Server{}
	err := srv.Start()
	if err != nil {
		log.Fatalf("Fail to start server, %q", err.Error())
	}
}
