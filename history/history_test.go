package history

import (
	"fmt"
	"os"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSave(t *testing.T) {
	h := New("history.json")
	h.Url.Add("http://localhost:8080/1")
	h.Url.Add("http://localhost:8080/2")
	h.Url.Add("http://localhost:8080/3")

	assert.Equal(t, 3, h.Url.Size())
	assert.Equal(t, "http://localhost:8080/3", h.Url.Get(h.Url.Size()-1))

	h.Url.Add("http://localhost:8080/4")
	assert.Equal(t, 4, h.Url.Size())

	h.Url.Add("http://localhost:8080/5")
	h.Url.Add("http://localhost:8080/6")
	h.Url.Add("http://localhost:8080/7")
	h.Url.Add("http://localhost:8080/8")
	h.Url.Add("http://localhost:8080/9")
	h.Url.Add("http://localhost:8080/10")
	h.Url.Add("http://localhost:8080/11")
	h.Url.Add("http://localhost:8080/12")

	h.Save()
}

func TestLoad(t *testing.T) {
	h := New("history.json")
	h.Load()
	fmt.Println(h)
}

func TestDeleteJson(t *testing.T) {
	os.Remove("history.json")
}
