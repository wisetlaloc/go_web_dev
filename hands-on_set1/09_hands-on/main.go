package main

import (
	"encoding/csv"
	"html/template"
	"log"
	"os"
	"strconv"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

type price struct {
	Date       string
	Open, High float64
}

func main() {
	csvFile, err := os.Open("./table.csv")
	check(err)
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	var prices []price
	data, error := reader.ReadAll()
	check(error)
	for i := 1; i < len(data); i++ {
		line := data[i]
		date := line[0]
		open, err := strconv.ParseFloat(line[1], 64)
		check(err)
		high, err := strconv.ParseFloat(line[2], 64)
		check(err)
		prices = append(prices, price{
			Date: date,
			Open: open,
			High: high,
		})
		//fmt.Println("aaa")
	}
	err = tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", prices)
	check(err)
}
