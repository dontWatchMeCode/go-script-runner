package main

import (
	"dontWatchMeCode/pipe/pkg/webhooks/discord"
)

func main() {
	webhookURL := "https://discord.com/api/webhooks/xyz123"

	message := discord.Message{
		Content: "example",
		Embeds: []discord.Embed{{
			Title:       "Embed Title",
			Description: "Embed Description",
			Color:       discord.EmbedsColorGreen,
		}},
	}

	discord.CallDiscordWebhook(webhookURL, message)
}
