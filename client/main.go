package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"

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

	accessCount int = 0
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

	fmt.Print("Enter server address: ")
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

	client := csmutex.NewCSMutexClient(conn)

	go mainLoop(client)

	fmt.Scanln()

}

func mainLoop(client csmutex.CSMutexClient) {
	for {
		logI("requesting token")
		_, err := client.RequestAccess(context.Background(), &csmutex.Identifier{
			Id: int32(os.Getpid()),
		})
		if err != nil {
			logE(err)
		}
		logI("token received, accessing critical section")

		_, err = client.PerformCriticalAction(context.Background(), &csmutex.ActionDetails{
			Id: &csmutex.Identifier{Id: int32(os.Getpid())},
		})
		if err != nil {
			logE(err)
		}
		accessCount++
		logI(fmt.Sprintf("critical section accessed %d times", accessCount))
		logI("critical section complete, releasing token")
		_, err = client.ReleaseAccess(context.Background(), &csmutex.Identifier{
			Id: int32(os.Getpid()),
		})
		if err != nil {
			logE(err)
		}
	}
}
