package server

import "net/http"

type Server struct {
	server *http.Server
}

func NewServer() *Server {
	svr := setupServer()
	return &Server{
		server: svr,
	}
}

func setupServer() *http.Server {
	mux := http.NewServeMux()
	svr := &http.Server{
		Handler: mux,
	}

	mux.Handle("/shorten", handlerShorten(nil))
	mux.Handle("/stats/", handlerStats())
	mux.Handle("/", handlerMain())

	return svr
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}
