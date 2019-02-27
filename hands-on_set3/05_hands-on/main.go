package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./tpl.gohtml"))
}

func main() {
	http.Handle("/", http.HandlerFunc(r))
	http.Handle("/dog/", http.HandlerFunc(d))
	http.Handle("/me/", http.HandlerFunc(m))
	http.ListenAndServe(":8080", nil)
}

func r(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "tpl.gohtml", "root")
	if err != nil {
		log.Fatalln(err)
	}
}

func d(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "tpl.gohtml", "dog")
	if err != nil {
		log.Fatalln(err)
	}
}

func m(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "tpl.gohtml", "wisetlaloc")
	if err != nil {
		log.Fatalln(err)
	}
}
