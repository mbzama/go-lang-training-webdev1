package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").ParseFiles("../global-functions/tpl.gohtml"))
}

type user struct {
	Name    string
	Message string
	Admin   bool
}

func main() {
	u1 := user{
		Name:    "user1",
		Message: "Message 1",
		Admin:   true,
	}

	u2 := user{
		Name:    "user2",
		Message: "Message 2",
		Admin:   false,
	}

	u3 := user{
		Name:    "",
		Message: "No Message",
		Admin:   false,
	}

	u4 := user{
		Name:    "user4",
		Message: "Message 2",
		Admin:   true,
	}

	users := []user{u1, u2, u3, u4}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", users)
	if err != nil {
		log.Fatalln(err)
	}
}
