package main

import (
	"dontWatchMeCode/pipe/pkg/webhooks/discord"
	"log"
)

func main() {

	data := discord.SendData{
		Url: "https://discord.com/api/webhooks/xyz123",
		Message: discord.Message{
			Content: "example",
			Embeds: []discord.Embed{{
				Title:       "Embed Title",
				Description: "Embed Description",
				Color:       discord.EmbedsColorGreen,
			}},
		},
	}

	err := discord.CallDiscordWebhook(data)
	if err != nil {
		log.Fatal(err)
	}
}
