package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").ParseFiles("../composition/tpl.gohtml"))
}

type person struct {
	Name string
	Age  int
}

func (p person) WelcomeMessage() string {
	return "Welcome " + p.Name + " !!!"
}

func (p person) AgeDouble() int {
	return p.Age * 2
}

func (p person) AddTenYears(x int) int {
	return x + 10
}

func main() {
	p := person{
		Name: "Alice",
		Age:  40,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", p)
	if err != nil {
		log.Fatalln(err)
	}
}
