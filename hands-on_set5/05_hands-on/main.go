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

func show(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/pics/", fs)
	http.HandleFunc("/", show)

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
