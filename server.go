package main

import (
	"fmt"
	"net/http"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) SetupRoutes() {
	s.mux = http.NewServeMux()
	s.mux.HandleFunc("/", s.Tabs)
}

func (s *Server) Tabs(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Number of open tabs: %d", 9999)
}

func (s *Server) Run() {
	fmt.Println("Listening on", *listenUrl)
	panic(http.ListenAndServe(*listenUrl, s.mux))
}
