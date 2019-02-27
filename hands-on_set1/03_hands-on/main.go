package main

import (
	"html/template"
	"os"
)

type hotel struct{ Name, Address, City, Zip, Region string }

type hotels []hotel

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	hotels := hotels{
		hotel{
			Name:    "Hotel A",
			Address: "511 A street",
			City:    "City A",
			Zip:     "123456",
			Region:  "southern",
		},
		hotel{
			Name:    "Hotel B",
			Address: "113 B street",
			City:    "City B",
			Zip:     "136789",
			Region:  "central",
		},
		hotel{
			Name:    "Hotel C",
			Address: "113 C street",
			City:    "City C",
			Zip:     "369467",
			Region:  "northern",
		},
	}

	tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", hotels)
}
