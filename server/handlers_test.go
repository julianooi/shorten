package server

import (
	"fmt"
	"github.com/julianooi/shorten"
	"github.com/julianooi/shorten/testhelper"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubShortener struct {
	called    bool
	toShorten string
}

func (s *StubShortener) Shorten(url string) (string, error) {
	s.called = true
	s.toShorten = url
	return "", nil
}

func TestHandlerShorten_Success(t *testing.T) {
	type request struct {
		URL string `json:"url"`
	}

	tt := []struct {
		url string
	}{
		{url: "http://digitalpenang.my"},
	}

	for _, tc := range tt {
		t.Run(tc.url, func(t *testing.T) {
			shortener := &StubShortener{}
			handler := handlerShorten(shortener)
			recorder := httptest.NewRecorder()
			rq := request{URL: tc.url}

			r := httptest.NewRequest(http.MethodPost, "/shorten", testhelper.MarshalJSONReader(t, rq))
			handler(recorder, r)

			if recorder.Code != http.StatusOK {
				t.Errorf("expected status to be %d, got %d", http.StatusOK, recorder.Code)
			}

			resp := map[string]string{}
			testhelper.UnmarshalJSON(t, recorder.Body, &resp)

			_, ok := resp["url"]
			if !ok {
				t.Errorf("expected url in response body json")
			}

			if !shortener.called {
				t.Errorf("expected shortener to be called")
			}

			if shortener.toShorten != tc.url {
				t.Errorf("shortener expected to be called with request body's url")
			}
		})
	}
}

type StubChecker struct {
	called bool
	key    string
}

func (s *StubChecker) Status(key string) (shorten.Status, error) {
	s.called = true
	s.key = key
	return shorten.Status{}, nil
}

func TestHandlerStats_Success(t *testing.T) {
	tt := []struct {
		key string
	}{
		{key: "abc123"},
	}

	for _, tc := range tt {
		t.Run(tc.key, func(t *testing.T) {
			checker := &StubChecker{}
			handler := handlerStats(checker)
			recorder := httptest.NewRecorder()
			rq := httptest.NewRequest("", fmt.Sprintf("/stats/%s", tc.key), nil)

			handler(recorder, rq)

			if recorder.Code != http.StatusOK {
				t.Errorf("expected code to be %d, got %d", http.StatusOK, recorder.Code)
			}

			if !checker.called {
				t.Errorf("checker expected to be called")
			}

			if checker.key != tc.key {
				t.Errorf("checker expected to be called with provided key [%s], got [%s]", tc.key, checker.key)
			}
		})
	}
}
