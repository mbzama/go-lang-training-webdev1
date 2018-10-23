package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	a := os.Args
	fmt.Println("No.of arguments: ", len(a))

	if len(a) < 2 {
		log.Fatal("Usage: go run ReadInputs.go {name}")
		os.Exit(1)
	}

	name := os.Args[1]

	str := fmt.Sprint(`
			<!DOCTYPE html>
		<html>
		<body>

		<h1 style="background-color:DodgerBlue;">Welcome ` + name + `</h1>

		</body>
		</html>

			`)

	nf, err := os.Create("index.html")

	if err != nil {
		log.Fatal("error creating file")
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))

	fmt.Println(nf)
}
