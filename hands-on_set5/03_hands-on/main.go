package main

import (
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {

	http.Handle("/", http.FileServer(http.Dir("starting-files")))

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
