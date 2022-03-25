package http

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCall(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "test")
	}))
	defer ts.Close()

	req := NewRequest("GET", ts.URL, nil, "")
	Call(req)
}
