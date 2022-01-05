package main

import (
	"io"
	"os"
	"example.com/sourse4"
)


func main() {
	myString := ""
	arguments := os.Args

	if len(arguments) == 1 {
		myString = "Plese, give me one argument !"
	} else {
		myString = arguments[1]
	}

	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")

	sourse4.Void()
}
