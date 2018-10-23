package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("../files/template2.gohtml", "../files/four.txt"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "template2.gohtml", "Md Zama")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "four.txt", "Md Zama")
	if err != nil {
		log.Fatalln(err)
	}

}
