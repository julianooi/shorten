package server

import "net/http"

type Server struct {
	server *http.Server
}

// Repository is a composite interface of all the functions
type Repository interface {
	Shortener
	StatsChecker
	StatsUpdater
}

// NewServer creates a new http server
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

// Start starts the http server
func (s *Server) Start() error {
	return s.server.ListenAndServe()
}
