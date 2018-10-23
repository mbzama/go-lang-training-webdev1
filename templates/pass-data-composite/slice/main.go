package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("template-slice.gohtml"))
}

func main() {
	//Passing Slice
	languages := []string{"C", "Java", "C#", "Python", "Go", "Node"}
	err := tpl.ExecuteTemplate(os.Stdout, "template-slice.gohtml", languages)
	if err != nil {
		log.Fatalln(err)
	}
}
