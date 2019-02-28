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
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", surf)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func surf(w http.ResponseWriter, _ *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
