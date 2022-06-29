package memory

import (
	"fmt"
	"testing"
)

func TestGenerateKey(t *testing.T) {
	tt := []struct {
		count    int
		expected string
	}{
		{
			count:    0,
			expected: "A-----",
		},
		{
			count:    10,
			expected: "K-----",
		},
		{
			count:    124,
			expected: "A99---",
		},
		{
			count:    371,
			expected: "999999",
		},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("count: %d", tc.count), func(t *testing.T) {
			key := generateKey(tc.count)
			if key != tc.expected {
				t.Errorf("expected [%s], got [%s]", tc.expected, key)
			}
		})
	}
}

func TestShortener_Shorten_Success(t *testing.T) {
	tt := []struct {
		url string
	}{
		{url: "https://www.example.com/long-test-url"},
	}
	for _, tc := range tt {
		t.Run(tc.url, func(t *testing.T) {
			shortener := NewShortener()
			key, err := shortener.Shorten(tc.url)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if len(key) != 6 {
				t.Errorf("unexpected key length, expected 6")
			}
		})
	}
}

func TestShortener(t *testing.T) {
	tt := []struct {
		url string
	}{
		{url: "https://www.google.com/my-long-test-string"},
	}

	for _, tc := range tt {
		t.Run(tc.url, func(t *testing.T) {
			shortener := NewShortener()
			key, err := shortener.Shorten(tc.url)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			status, err := shortener.Status(key)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if status.Count != 0 {
				t.Errorf("expected count to be 0, got %d", status.Count)
			}

			err = shortener.Increment(key)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			status, err = shortener.Status(key)
			if status.Count != 1 {
				t.Errorf("expected count to be 1, got %d", status.Count)
			}
		})
	}
}
