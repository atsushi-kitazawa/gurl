package http

import (
	"os"
	"testing"
)

func httpServer() {

}

func TestMain(m *testing.M) {
	code := m.Run()

	os.Exit(code)
}

func TestCall(t *testing.T) {
	req := NewRequest("GET", "http://localhost:8080", nil, "body data")
	Call(req)
}
