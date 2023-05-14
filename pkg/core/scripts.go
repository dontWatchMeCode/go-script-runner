package core

import (
	"dontWatchMeCode/pipe/pkg/webhooks/discord"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func RunAllScript(pwd string) {
	fmt.Println("Running all scripts")
	var embeds []discord.Embed
	date := time.Now().Format("2006-01-02 15:04:05")

	defer func() {
		discordWebhookUrl, err := GetEnv("DISCORD_WEBHOOK_URL")
		if err != nil {
			panic(err)
		}

		var filteredEmbeds []discord.Embed
		for _, embed := range embeds {
			if embed.Description != "" {
				filteredEmbeds = append(filteredEmbeds, embed)
			}
		}

		fmt.Println("Sending discord webhook")
		discord.CallDiscordWebhook(discord.SendData{
			Url: discordWebhookUrl,
			Message: discord.Message{
				Content: "",
				Embeds:  filteredEmbeds,
			},
		})
	}()

	logFile := getFileReference(filepath.Join(pwd, "scripts.log"))
	defer logFile.Close()

	if err := os.Chdir(filepath.Join(pwd, "scripts")); err != nil {
		embeds = append(embeds, discord.Embed{
			Title:       fmt.Sprintf("[ WARNING: %s / %s ]", "folder scripts", date),
			Description: "folder scripts not found,\nplease create folder scripts",
			Color:       discord.EmbedsColorYellow,
		})

		index := len(embeds) - 1

		LogData(
			logFile,
			embeds[index].Title,
			embeds[index].Description,
		)

		panic(err)
	}

	files, err := os.ReadDir(".")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() || strings.HasPrefix(file.Name(), "_") {
			continue
		}

		cmd := exec.Command(
			"bash", "-e",
			filepath.Join(pwd, "scripts", file.Name()),
		)
		out, err := cmd.Output()

		if err != nil {
			embeds = append(embeds, discord.Embed{
				Title:       fmt.Sprintf("[ ERROR: %s / %s ]", file.Name(), date),
				Description: strings.TrimSpace(string(err.Error())),
				Color:       discord.EmbedsColorYellow,
			})

			index := len(embeds) - 1
			title := embeds[index].Title
			description := embeds[index].Description
			if description == "" {
				description = "no output"
			}

			LogData(
				logFile,
				title,
				embeds[index].Description,
			)
			continue
		}

		embeds = append(embeds, discord.Embed{
			Title:       fmt.Sprintf("[ INFO: %s / %s ]", file.Name(), date),
			Description: strings.TrimSpace(string(out)),
			Color:       discord.EmbedsColorGreen,
		})

		index := len(embeds) - 1
		title := embeds[index].Title
		description := embeds[index].Description
		if description == "" {
			description = "no output"
		}

		LogData(logFile, title, description)
	}
}
