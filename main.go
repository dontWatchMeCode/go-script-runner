package main

import (
	"dontWatchMeCode/pipe/pkg/utils"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {
	defer utils.HandlePanic()
	godotenv.Load(".env")

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	utils.StartCron(pwd)

	fmt.Println("Press CTRL-C to exit")
	<-signalChannel
	fmt.Println("Program exiting")
}
