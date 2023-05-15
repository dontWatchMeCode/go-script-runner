package main

import (
	"dontWatchMeCode/pipe/pkg/core"
	"dontWatchMeCode/pipe/pkg/webhooks/discord"
	"fmt"
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
	} else if len(os.Args[1:]) != 0 && os.Args[1] == "-test" {
		fmt.Println("Sending discord webhook test")
		discordWebhookUrl, err := core.GetEnv("DISCORD_WEBHOOK_URL")
		if err != nil {
			panic(err)
		}
		discord.CallDiscordWebhookTest(discordWebhookUrl)
	} else {
		signalChannel := make(chan os.Signal, 1)
		signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

		core.StartCron(pwd)

		<-signalChannel
	}
}
