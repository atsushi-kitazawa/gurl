package history

import (
	"fmt"
	"testing"
)

func TestURLHistory(t *testing.T) {
	uh := New()
	uh.Push("aaa")
	uh.Push("bbb")
	uh.Push("ccc")
	fmt.Println(uh)       // aaa,bbb,ccc
	fmt.Println(uh.Pop()) //ccc
	fmt.Println(uh)       // aaa,bbb
	uh.Push("ddd")
	fmt.Println(uh.Pop()) //ddd
	fmt.Println(uh.Pop()) //bbb
	fmt.Println(uh)       // aaa
}
