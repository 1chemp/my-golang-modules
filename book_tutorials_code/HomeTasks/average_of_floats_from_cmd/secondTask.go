package main

/*
	Напишите Go-программу, вычисляющую среднее значение всех чисел с пла-
	вающей запятой, переданных программе в качестве аргументов командной
	строки.
*/

import (
	"fmt"
	"os"
	"errors"
	"strconv"
)


func main() {

	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Pls give me one or more floats numbers.")
		os.Exit(1)
	}

	k := 1
	var err error = errors.New("Error !")
	var n float64 = 0

	for err != nil {
		if k >= len(arguments) {
			fmt.Println("None of this arguments of cmd is float number.")
			return
		}
		n, err = strconv.ParseFloat(arguments[k], 64)
		k++
	}
	summary := n
	var lenght float64 = 1

	for i:=2; i<len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)

		if err == nil {
			summary += n
			lenght++
		}
	}

	average := summary/lenght
	fmt.Println("Summary: ", summary)
	fmt.Println("Lenght: ", lenght)
	fmt.Println("Average of floats numbers: ", average)
	return
}


