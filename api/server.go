package api

import (
	"net/http"
)

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) Server {

	return Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/api/v1/codes", s.HandleAllCodes)
	http.HandleFunc("/api/v1/codes/{code}", s.HandleCode)
	http.HandleFunc("/api/v1/codes/parse", s.HandleParse)

	return http.ListenAndServe(s.listenAddr, nil)
}
