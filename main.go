package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/atsushi-kitazawa/gurl/http"
)

func main() {
	doMain()
}

func doMain() {
	fmt.Printf("URL : ")
	url := readKeyboard()
	fmt.Printf("Method : ")
	method := readKeyboard()
	fmt.Printf("Headers : ")
	_ = readKeyboard()
	fmt.Printf("Body : ")
	body := readKeyboard()

	req := http.NewRequest(method, url, nil, body)
	http.Call(req)
}

func readKeyboard() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input[:len(input)-1]
}
