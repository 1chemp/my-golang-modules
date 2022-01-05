package main

/*
	Вызов и обработка вызываемой ошибки
	посредством средств логирования
*/

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var LOGFILE = "/tmp/mGo.log"

func returnError(a, b int) error {
	if a == b {
		err := errors.New("Error in returnError() function!")
		return err
	} else {
		return nil
	}
}


func main() {
	err := returnError(1, 2)
	f, _ := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	iLog := log.New(f, "ErrorInCompare", log.LstdFlags)
	iLog.SetFlags(log.LstdFlags)

	if err == nil {
		fmt.Println("returnError() ended normally!")
	} else {
		iLog.Println("Hello there!")
		iLog.Println("Another log entry!")
	}

	err = returnError(10, 10)
	if err == nil {
		fmt.Println("returnError() ended normally!")
	} else {
		iLog.Println("Hello there!")
		iLog.Println("Another log entry!")

	}

	if err.Error() == "Error in returnError() function!" {
		fmt.Println("Log in the stdout ??!!")
	}

}
