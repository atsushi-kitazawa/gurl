package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/atsushi-kitazawa/gurl/history"
	"github.com/atsushi-kitazawa/gurl/http"
	"github.com/atsushi-kitazawa/gurl/print"
)

func main() {
	doMain()
}

func doMain() {
	history := history.New("history.json")
	history.Load()

	fmt.Printf("URL : ")
	url := readKeyboard()
	fmt.Printf("Method : ")
	method := readKeyboard()
	fmt.Printf("Headers : ")
	headers := readKeyboard()
	fmt.Printf("Body : ")
	body := readKeyboard()

	req := http.NewRequest(method, url, http.CreateHeader(headers), body)
	res := http.Call(req)
	print.Print(res)

	history.Url.Add(url)
	history.Save()
}

func readKeyboard() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input[:len(input)-1]
}
