package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("../files/template1.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "template1.gohtml", "Zama")
	if err != nil {
		log.Fatalln(err)
	}
}
