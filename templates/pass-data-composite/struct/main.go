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
	tpl = template.Must(template.ParseFiles("template-struct.gohtml"))
}

func main() {
	//Passing Slice
	l1 := language{
		Order: "1",
		Name:  "C",
	}
	err := tpl.ExecuteTemplate(os.Stdout, "template-struct.gohtml", l1)
	if err != nil {
		log.Fatalln(err)
	}
}
