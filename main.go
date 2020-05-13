package main

import (
	"log"
	"net/http"

	"github.com/egawata/gohello/getname"
)

type MyHandler struct{}

func (s *MyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	content := []byte(`Hello, ` + getname.Name())
	res.Write(content)
}

func main() {
	h := &MyHandler{}
	http.Handle("/", h)
	log.Fatal(http.ListenAndServe(":8088", nil))
}
