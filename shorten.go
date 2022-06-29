package shorten

import (
	"errors"
	"time"
)

// Request is the form used for shortening the url
type Request struct {
	URL string `json:"url"`
}

// Status consists of all the status information
type Status struct {
	URL       string    `json:"url"`
	Count     int       `json:"count"`
	CreatedAt time.Time `json:"createdAt"`
}

// ErrKeyNotFound indicates that the key specified is not found
var ErrKeyNotFound = errors.New("key not found")
