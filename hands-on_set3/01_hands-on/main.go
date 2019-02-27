package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", r)
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/me/", m)
	http.ListenAndServe(":8080", nil)
}

func r(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "I'm in root")
}

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "I'm in dog")
}

func m(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "I'm wisetlaloc")
}
