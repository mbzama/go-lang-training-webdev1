package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type language struct {
	Order string
	Name  string
}

func init() {
	tpl = template.Must(template.ParseFiles("template-slice-struct.gohtml"))
}

func main() {
	//Passing Slice of struct
	languages := []language{
		{
			Order: "1",
			Name:  "C",
		},
		{
			Order: "2",
			Name:  "Java",
		},
		{
			Order: "3",
			Name:  "Python",
		},
		{
			Order: "4",
			Name:  "Go",
		},
	}
	err := tpl.ExecuteTemplate(os.Stdout, "template-slice-struct.gohtml", languages)
	if err != nil {
		log.Fatalln(err)
	}
}
