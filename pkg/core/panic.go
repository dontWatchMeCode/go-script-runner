package core

import (
	"dontWatchMeCode/pipe/pkg/webhooks/discord"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func HandlePanic() {
	if r := recover(); r != nil {
		embed := discord.Embed{
			Title:       fmt.Sprintf("[ FATAL: %s ]", "Panic occurred"),
			Description: fmt.Sprintf("%v", r),
			Color:       discord.EmbedsColorRed,
		}

		discordWebhookUrl, err := GetEnv("DISCORD_WEBHOOK_URL")
		if err != nil {
			panic(err)
		}

		fmt.Println("Sending discord webhook")
		discord.CallDiscordWebhook(discord.SendData{
			Url: discordWebhookUrl,
			Message: discord.Message{
				Content: "",
				Embeds:  []discord.Embed{embed},
			},
		})

		pwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		logFile := getFileReference(filepath.Join(pwd, "scripts.log"))
		defer logFile.Close()

		LogData(
			logFile,
			embed.Title,
			embed.Description,
		)

		log.Printf("Panic occurred: %v", r)
		os.Exit(1)
	}
}
