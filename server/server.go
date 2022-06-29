package server

import "net/http"

type Server struct {
	server *http.Server
}

type Repository interface {
	Shortener
	StatsChecker
	StatsUpdater
}

func NewServer(repository Repository) *Server {
	svr := setupServer(repository)
	return &Server{
		server: svr,
	}
}

func setupServer(repository Repository) *http.Server {
	mux := http.NewServeMux()
	svr := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	mux.Handle("/shorten", handlerShorten(repository))
	mux.Handle("/stats/", handlerStats(repository))
	mux.Handle("/", handlerMain(repository, repository))

	return svr
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}
