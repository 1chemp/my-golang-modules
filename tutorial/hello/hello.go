package main

import (
    "fmt"
    "log"
    "example.com/greetings"
)

//import "rsc.io/quote"

func main() {
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    names := []string{"Gladys", "Samantha", "Darrin"}
    message, err := greetings.Hellos(names)
    //fmt.Println(message)
    if err != nil {
	log.Fatal(err)
    }
    fmt.Println(message)
}

