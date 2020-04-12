package server

import (
	"fmt"
	"net/http"
)

type Server struct {
}

// do the server struct to implement the http/Handler
func (s Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	_, _ = fmt.Fprintf(writer, "ok")
}

func (s Server) Start() error {
	return http.ListenAndServe(":8080", s)
}