package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpltext *template.Template
var tplhtml *template.Template

func init() {
	tpltext = template.Must(template.ParseGlob("../files/*.txt"))
	tplhtml = template.Must(template.ParseGlob("../files/*.gohtml"))
}

func main() {
	fmt.Println("---------")
	fmt.Println("HTML")
	fmt.Println("---------")
	err := tplhtml.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("---------")
	err = tplhtml.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tplhtml.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("---------")
	err = tplhtml.ExecuteTemplate(os.Stdout, "three.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("---------")
	fmt.Println("Text")
	fmt.Println("---------")
	err = tpltext.ExecuteTemplate(os.Stdout, "five.txt", nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("---------")
	err = tpltext.ExecuteTemplate(os.Stdout, "six.txt", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
