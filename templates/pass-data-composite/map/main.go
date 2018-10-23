package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("template-map.gohtml"))
}

func main() {
	//Passing Map
	languages := map[string]string{
		"1": "C",
		"2": "Java",
		"3": "C#",
		"4": "Python",
		"5": "Go",
	}
	err := tpl.ExecuteTemplate(os.Stdout, "template-map.gohtml", languages)
	if err != nil {
		log.Fatalln(err)
	}
}
