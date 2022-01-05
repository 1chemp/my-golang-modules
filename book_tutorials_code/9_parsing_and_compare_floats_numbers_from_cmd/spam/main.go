package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Please, give one or more floats numbers.")
		os.Exit(1)
	}

	arguments := os.Args
	var err error = errors.New("An error")
	k := 1
	var n float64

	for err != nil {
		if k >= len(arguments) {
			fmt.Println("None of the arguments is a float!")
			return
		}
		n, err = strconv.ParseFloat(arguments[k], 64)
		k++
	}
	min, max := n, n
	fmt.Println("Start max value: ", max)
	fmt.Println("Start min value: ", min)

	for i := 2; i < len(arguments); i++ {
		n, err = strconv.ParseFloat(arguments[i], 64)
		if err == nil {
			if n < min {
				min = n
			}
			if n > max {
				max = n
			}
		}
	}

	fmt.Println("Finish max value: ", max)
	fmt.Println("Finish min value: ", min)
}
