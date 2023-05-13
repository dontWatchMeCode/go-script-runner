package main

import (
	"dontWatchMeCode/pipe/pkg/core"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {
	defer core.HandlePanic()
	godotenv.Load(".env")

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if len(os.Args[1:]) != 0 && os.Args[1] == "-run" {
		core.RunAllScript(pwd)
	} else {
		signalChannel := make(chan os.Signal, 1)
		signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

		core.StartCron(pwd)

		<-signalChannel
	}
}
