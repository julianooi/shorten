package server

import "net/http"

// handlerShorten is a http handler to shorten a url
func handlerShorten() http.HandlerFunc {
	type request struct {
		URL string `json:"url"`
	}

	return func(writer http.ResponseWriter, request *http.Request) {
		// TODO: shorten url, store into memory, return new url
	}
}

func handlerStats() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// TODO: retrieve path from request
		// TODO: retrieve stats from memory
		// TODO: return stats
	}
}

func handlerMain() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// TODO: retrieve path from request
		// TODO: redirect to url
	}
}
