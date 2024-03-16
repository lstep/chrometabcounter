package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Server struct {
	mux *http.ServeMux
}

type TabToReturn struct {
	Description string `json:"description"`
	Count       int    `json:"count"`
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) SetupRoutes() {
	s.mux = http.NewServeMux()
	s.mux.HandleFunc("/", s.Tabs)
}

func (s *Server) Tabs(w http.ResponseWriter, r *http.Request) {
	numberOfTtabs, err := CountTabs()
	if err != nil {
		fmt.Println("Error retrieving tabs:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := TabToReturn{
		Description: "Number of open tabs",
		Count:       numberOfTtabs,
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func (s *Server) Run() {
	fmt.Println("Listening on", *listenUrl)
	panic(http.ListenAndServe(*listenUrl, s.mux))
}
