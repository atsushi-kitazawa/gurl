package print

import (
	"bytes"
	"fmt"
	"net/http"
)

func Print(r *http.Response) {
	fmt.Println("===== response info =====")

	fmt.Printf("Status : %v \n", r.Status)

	for k, v := range r.Header {
		fmt.Printf("Headers : %v:%v \n", k, v)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	fmt.Printf("Body : %v \n", buf)
}
