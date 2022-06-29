package memory

import (
	"github.com/julianooi/shorten"
	"sync"
	"time"
)

type status struct {
	url       string
	createdAt time.Time
	mu        sync.Mutex
	count     int
}

func (s *status) increment() {
	s.mu.Lock()
	s.count++
	s.mu.Unlock()
}

type Shortener struct {
	store map[string]*status
	mu    sync.Mutex
	count int
}

func NewShortener() *Shortener {
	return &Shortener{
		store: make(map[string]*status),
	}
}

func (s *Shortener) nextKey() string {
	s.mu.Lock()
	c := s.count
	s.count++
	s.mu.Unlock()
	return generateKey(c)
}

var validChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func generateKey(count int) string {
	l := len(validChars)
	output := ""

	output += string(validChars[count%l])
	prependCount := count / l
	for i := 0; i < prependCount; i++ {
		output += string(validChars[l-1])
	}

	toAppend := 6 - len(output)
	if toAppend < 0 {
		panic("more characters than support[6]")
	}

	for i := 0; i < toAppend; i++ {
		output += "-"
	}

	return output
}

func (s *Shortener) Shorten(url string) (string, error) {
	key := s.nextKey()
	s.store[key] = &status{
		url:       url,
		createdAt: time.Now(),
		mu:        sync.Mutex{},
		count:     0,
	}
	return key, nil
}

func (s *Shortener) Status(key string) (shorten.Status, error) {
	state, ok := s.store[key]
	if !ok {
		return shorten.Status{}, shorten.ErrKeyNotFound
	}
	return shorten.Status{
		URL:       state.url,
		Count:     state.count,
		CreatedAt: state.createdAt,
	}, nil
}

func (s *Shortener) Increment(key string) error {
	state, ok := s.store[key]
	if !ok {
		return shorten.ErrKeyNotFound
	}
	state.increment()
	return nil
}
