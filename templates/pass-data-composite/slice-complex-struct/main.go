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

type ide struct {
	Name   string
	Rating int
}

type summary struct {
	Languages []language
	Ides      []ide
}

func init() {
	tpl = template.Must(template.ParseFiles("template-slice-struct.gohtml"))
}

func main() {
	//Passing Slice of struct
	l := []language{
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

	i := []ide{
		{Name: "Visual Studio Code", Rating: 1},
		{Name: "Atom", Rating: 2},
		{Name: "RStudio", Rating: 3},
		{Name: "Eclipse", Rating: 4},
	}

	s := summary{
		Languages: l, Ides: i,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "template-slice-struct.gohtml", s)
	if err != nil {
		log.Fatalln(err)
	}
}
