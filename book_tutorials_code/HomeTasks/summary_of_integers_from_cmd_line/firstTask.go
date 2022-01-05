package main

/*

	Напишите Go-программу, 
	которая вычисляет сумму всех аргументов командной
	строки, которые являются действительными числами.

*/

import (
	"fmt"
	"strconv"
	"os"
	"errors"
)


func main() {

	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Pls give me one or more real numbers.")
		os.Exit(1)
	}

	var err error = errors.New("Error!")
	n := 0
	k := 1

	for err != nil {
		if k >= len(arguments) {
			fmt.Println("None of the arguments is a real Number")
			return
		}
		n, err = strconv.Atoi(arguments[k])
		k++
	}

	summary := n

	for i := 2; i < len(arguments); i++ {
		n, err := strconv.Atoi(arguments[i])
		if err == nil && n >= 1 {
			summary += n
		}
	}

	fmt.Println("Summary of real numbers: ", summary)
}
