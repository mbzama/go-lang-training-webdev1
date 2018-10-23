package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

//Create a funcMap to register functions
//"uc" is what the func will be called in the template, calls ToUpper func from package strings
//"ft" is a func declared in this file, slices a string and returns first 3 characters
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(i string) string {
	s := strings.TrimSpace(i)
	if len(s) > 3 {
		return s[:3]
	}
	return s
}

type language struct {
	Order string
	Name  string
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("template-functions.gohtml"))
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
	err := tpl.ExecuteTemplate(os.Stdout, "template-functions.gohtml", languages)
	if err != nil {
		log.Fatalln(err)
	}
}
