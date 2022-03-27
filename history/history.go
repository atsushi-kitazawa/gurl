package history

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	HistoryMaxSize = 10
)

type URLHistory []string

type History struct {
	file string
	Url  URLHistory `json:"url"`
}

func New(file string) *History {
	return &History{
		file: file,
		Url:  make([]string, 0),
	}
}

func (u *URLHistory) Add(e string) {
	uu := *u
	if u.Size() >= HistoryMaxSize {
		uu = uu[1:]
	}
	*u = append(uu, e)
}

func (u *URLHistory) Get(index int) string {
	uu := *u
	return uu[index]
}

func (u *URLHistory) Size() int {
	uu := *u
	return len(uu)
}

func (h *History) Load() {
	data, err := ioutil.ReadFile(h.file)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(data, h); err != nil {
		panic(err)
	}
}

func (h *History) Save() {
	data, err := json.MarshalIndent(h, "", "\t")
	if err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile(h.file, data, os.ModePerm); err != nil {
		panic(err)
	}
}

func (u *URLHistory) String() string {
	return fmt.Sprint(*u)
}
