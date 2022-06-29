package server

import (
	"github.com/julianooi/shorten/testhelper"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubShortener struct {
	called bool
}

func (s *StubShortener) Shorten(url string) (string, error) {
	s.called = true
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

			r := httptest.NewRequest("", "/shorten", testhelper.MarshalJSONReader(t, rq))
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
		})
	}
}
