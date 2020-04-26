package server

import (
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
	ch chan os.Signal
}

// do the server struct to implement the http/Handler
func (s Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	_, _ = fmt.Fprintf(writer, "ok")
}

func (s *Server) quit() {
	s.ch <- quitSignal
}

func (s *Server) Start() error {
	var err error = nil
	signal.Notify(s.ch, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Print("Starting server in port 8080 ...")
		if err = http.ListenAndServe(":8080", s); err != nil {
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
	}

	return err
}

func NewServer() *Server {
	srv := Server{
		ch: make(chan os.Signal, 1),
	}

	return &srv
}