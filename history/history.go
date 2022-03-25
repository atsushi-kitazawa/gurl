package history

import "fmt"

type History interface {
	Push(string)
	Pop() string
	Save(*History)
}

type URLHistory struct {
	entry []string
}

func New() History {
	return &URLHistory{
		entry: make([]string, 0),
	}
}

func (u *URLHistory) Push(e string) {
	u.entry = append(u.entry, e)
}

func (u *URLHistory) Pop() string {
	ret := u.entry[len(u.entry)-1]
	u.entry = u.entry[:len(u.entry)-1]
	return ret
}

func (u *URLHistory) Save(h *History) {
}

func (u *URLHistory) String() string {
	return fmt.Sprint(u.entry)
}
