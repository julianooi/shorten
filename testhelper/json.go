package testhelper

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"
)

func MarshalJSONReader(t *testing.T, target any) io.Reader {
	t.Helper()
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	err := enc.Encode(target)
	if err != nil {
		t.Fatal(err)
	}
	return buf
}

func UnmarshalJSON(t *testing.T, reader io.Reader, target any) {
	t.Helper()
	dec := json.NewDecoder(reader)
	err := dec.Decode(target)
	if err != nil {
		t.Fatal(err)
	}
}
