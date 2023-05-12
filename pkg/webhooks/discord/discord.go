package discord

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	EmbedsColorGreen  = 65280    // 0x00FF00
	EmbedsColorYellow = 16776960 // 0xFFFF00
	EmbedsColorRed    = 16711680 // 0xFF0000
)

type Embed struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Color       int    `json:"color,omitempty"`
}

type Message struct {
	Content string  `json:"content,omitempty"`
	Embeds  []Embed `json:"embeds,omitempty"`
}

func CallDiscordWebhook(webhookURL string, message Message) {
	payload := Message{
		Content: message.Content,
		Embeds:  message.Embeds,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatal("Failed to marshal JSON payload:", err)
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewReader(payloadBytes))
	if err != nil {
		log.Fatal("Failed to send HTTP request:", err)
	}
	defer resp.Body.Close()

	if !strings.HasPrefix(strconv.Itoa(resp.StatusCode), "2") {
		log.Println("Unexpected response status:", resp.StatusCode)
	} else {
		log.Println("Discord webhook called successfully")
	}
}
