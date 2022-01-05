package main

/*
Напишите Go-программу, которая считывает целые числа до тех пор, пока
не встретит во входных данных слово END .
*/


import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var f *os.File
	f = os.Stdin

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() != "END" {
			fmt.Println(">", scanner.Text())
		} else {
			fmt.Println("Exiting...")
			break
		}
	}
	return

}
