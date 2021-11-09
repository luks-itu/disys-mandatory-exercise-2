package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/luks-itu/disys-mandatory-exercise-2/csmutex"
	"google.golang.org/grpc"
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
	defer fmt.Println(&buf)
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

	fmt.Println("Enter server address: ")
	fmt.Scanln(&hostAddress)

	startClient(hostAddress)

}

func startClient(address string) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		logF(err.Error())
	}

	defer conn.Close()

	client := csmutex.NewCSMutexClient
}
