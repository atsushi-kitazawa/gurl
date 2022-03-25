package http

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCall(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "test")
	}))
	defer ts.Close()

	req := NewRequest("GET", ts.URL, nil, "")
	Call(req)
}

func TestCreateHeader(t *testing.T) {
	data := "aaa:AAA,bbb: BBB, ccc : CCC"
	r := CreateHeader(data)
	assert.Equal(t, r["aaa"], "AAA")
	assert.Equal(t, r["bbb"], "BBB")
	assert.Equal(t, r["ccc"], "CCC")
}
