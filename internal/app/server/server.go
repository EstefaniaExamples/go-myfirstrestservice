package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	quitSignal = syscall.SIGQUIT
)

type Server struct {
	http.Server
	ch   chan os.Signal
	port int
}

// do the server struct to implement the http/Handler
func (s Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	_, _ = fmt.Fprintf(writer, "ok")
}

func (s *Server) quit() {
	log.Print("Calling quit signal ... ")
	s.ch <- quitSignal
}

func (s *Server) Start() error {
	var err error = nil
	signal.Notify(s.ch, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("Starting server in port %d ...", s.port)
		if err = s.ListenAndServe(); err != nil {
			log.Print("Error during listen and server ...")
			s.quit()
		}
	}()

	if err == nil {
		killSignal := <-s.ch
		switch killSignal {
		case os.Interrupt:
			log.Print("Got interrupt signal closing ...")
		case syscall.SIGTERM:
			log.Print("Got termination signal closing ...")
		case quitSignal:
			log.Print("Got quit signal closing ...")
		}

		_ = s.Shutdown(context.Background())
	}

	return err
}

func NewServer(port int) *Server {
	strPort := fmt.Sprintf(":%d", port)
	srv := Server{
		Server: http.Server{
			Addr: strPort,
			// I am the handler cause I am embedding from http.Server
		},
		ch:   make(chan os.Signal, 1),
		port: port,
	}

	return &srv
}
