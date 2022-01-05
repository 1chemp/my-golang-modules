package sourse4

import (
	"bufio"
	"fmt"
	"os"
)

func Void() {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
	fmt.Println(">", scanner.Text())
	}
}

