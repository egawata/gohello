package main

import (
	"log"
	"net/http"
)

type MyHandler struct{}

func (s *MyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	content := []byte(`Welcome to github.com/egawata/gohello`)
	res.Write(content)
}

func main() {
	h := &MyHandler{}
	http.Handle("/", h)
	log.Fatal(http.ListenAndServe(":8088", nil))
}
