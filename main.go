package main

import (
	"dontWatchMeCode/pipe/pkg/utils"
	"dontWatchMeCode/pipe/pkg/webhooks/discord"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

func getFileReference(filePath string) *os.File {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		logFile, err := os.Create(filePath)
		if err != nil {
			panic(err)
		}
		defer logFile.Close()
	}

	fileReference, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	return fileReference
}

func logData(file *os.File, heading string, content string) {
	splitString := strings.Repeat("-", len(heading))
	fmt.Fprintf(file, "%s\n%s\n%s\n%s\n\n", splitString, heading, splitString, content)
}

func handlePanic() {
	if r := recover(); r != nil {
		embed := discord.Embed{
			Title:       fmt.Sprintf("[ FATAL: %s ]", "Panic occurred"),
			Description: fmt.Sprintf("%v", r),
			Color:       discord.EmbedsColorRed,
		}

		discordWebhookUrl, error := utils.GetEnv("DISCORD_WEBHOOK_URL")
		if error != nil {
			panic(error)
		}

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

		logData(
			logFile,
			embed.Title,
			embed.Description,
		)

		log.Printf("Panic occurred: %v", r)
		os.Exit(1)
	}
}

func runAllScript() {
	var embeds []discord.Embed

	defer func() {
		discordWebhookUrl, error := utils.GetEnv("DISCORD_WEBHOOK_URL")
		if error != nil {
			panic(error)
		}

		discord.CallDiscordWebhook(discord.SendData{
			Url: discordWebhookUrl,
			Message: discord.Message{
				Content: "",
				Embeds:  embeds,
			},
		})
	}()

	date := time.Now().Format("2006-01-02 15:04:05")
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	logFile := getFileReference(filepath.Join(pwd, "scripts.log"))
	defer logFile.Close()

	if err := os.Chdir(filepath.Join(pwd, "scripts")); err != nil {
		embeds = append(embeds, discord.Embed{
			Title:       fmt.Sprintf("[ WARNING: %s / %s ]", "folder scripts", date),
			Description: "folder scripts not found,\nplease create folder scripts",
			Color:       discord.EmbedsColorYellow,
		})

		index := len(embeds) - 1

		logData(
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

			logData(
				logFile,
				embeds[index].Title,
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

		logData(
			logFile,
			embeds[index].Title,
			embeds[index].Description,
		)
	}
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	defer handlePanic()
	runAllScript()
}
