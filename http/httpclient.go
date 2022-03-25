package http

import (
	"net/http"
)

type Request struct {
	method  string
	url     string
	headers map[string]string
	body    string
}

func NewRequest(method string, url string, headers map[string]string, body string) *Request {
	return &Request{
		method:  method,
		url:     url,
		headers: headers,
		body:    body,
	}
}

func Call(req *Request) *http.Response {
	var res *http.Response
	var err error
	switch req.method {
	case "GET":
		res, err = http.Get(req.url)
	case "POST":
		res, err = http.Post(req.url, "", nil)
	}
	if err != nil {
		panic(err)
	}

	return res
}
