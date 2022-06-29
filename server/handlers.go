package server

import (
	"encoding/json"
	"fmt"
	"github.com/julianooi/shorten"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Shortener interface {
	Shorten(url string) (string, error)
}

// handlerShorten is a http handler to shorten a url
func handlerShorten(shortener Shortener) http.HandlerFunc {
	type request struct {
		URL string `json:"url"`
	}

	type response struct {
		URL string `json:"url"`
	}

	return func(w http.ResponseWriter, rq *http.Request) {
		r := request{}
		body := rq.Body
		defer func() {
			_ = body.Close()
		}()
		dec := json.NewDecoder(body)
		err := dec.Decode(&r)
		if err != nil {
			http.Error(w, "failed to decode body", http.StatusInternalServerError)
			return
		}

		_, err = url.Parse(r.URL)
		if err != nil {
			http.Error(w, "invalid url in request", http.StatusBadRequest)
			return
		}

		shortened, err := shortener.Shorten(r.URL)
		if err != nil {
			http.Error(w, "failed to shorten url", http.StatusInternalServerError)
			log.Println(fmt.Errorf("failed to shorten url: %w", err))
			return
		}

		w.WriteHeader(http.StatusOK)
		enc := json.NewEncoder(w)
		err = enc.Encode(response{URL: shortened})
		if err != nil {
			log.Println(fmt.Errorf("failed to encode response: %w", err))
		}
	}
}

type StatsChecker interface {
	Status(key string) (shorten.Status, error)
}

func handlerStats(checker StatsChecker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := r.URL.Path
		parts := strings.Split(p, "/")

		key := parts[len(parts)-1]

		status, err := checker.Status(key)
		if err != nil {
			http.Error(w, "failed to check status", http.StatusInternalServerError)
			log.Println(fmt.Errorf("failed to check status: %w", err))
			return
		}

		enc := json.NewEncoder(w)
		err = enc.Encode(status)
		if err != nil {
			log.Println(fmt.Errorf("failed to encode status: %w", err))
		}
	}
}

func handlerMain() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// TODO: retrieve path from request
		// TODO: redirect to url
	}
}
