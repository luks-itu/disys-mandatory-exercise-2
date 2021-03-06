package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

var (
	hostAddress = "localhost:32610"

	// Logs str
	logI func(str string)

	// Logs str as an error
	logE func(err error)

	// Logs str as a fatal error, then panics
	logF func(str string)
)

func main() {
	var buf bytes.Buffer
	var loggerInf = log.New(&buf, "LOG|INFO: ", log.Lshortfile|log.Lmicroseconds)
	var loggerErr = log.New(&buf, "LOG|ERR: ", log.Lshortfile|log.Lmicroseconds)
	var loggerFat = log.New(&buf, "LOG|FATAL: ", log.Lshortfile|log.Lmicroseconds)
	defer func() {
		fmt.Println(&buf)
		os.WriteFile("log.txt", buf.Bytes(), 0644)
	}()
	logI = func(str string) {
		loggerInf.Output(2, str)
		fmt.Println(str)
	}
	logE = func(err error) {
		loggerErr.Output(2, err.Error())
		fmt.Println(err.Error())
	}
	logF = func(str string) {
		loggerFat.Output(2, str)
		panic(str)
	}

	fmt.Print("Enter hostAddress: ")
	fmt.Scanln(&hostAddress)

	go startServer(hostAddress)

	fmt.Scanln()
}
