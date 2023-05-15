package discord

import (
	"bytes"
	"encoding/json"
	"errors"
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

type SendData struct {
	Url     string
	Message Message
}

func CallDiscordWebhook(data SendData) error {
	payload := Message{
		Content: data.Message.Content,
		Embeds:  data.Message.Embeds,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(data.Url, "application/json", bytes.NewReader(payloadBytes))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if strings.HasPrefix(strconv.Itoa(resp.StatusCode), "2") {
		return nil
	}

	return errors.New("Webhook request failed status:" + strconv.Itoa(resp.StatusCode))
}

func CallDiscordWebhookTest(url string) error {
	return CallDiscordWebhook(
		SendData{
			Url: url,
			Message: Message{
				Content: "",
				Embeds: []Embed{
					{
						Title:       "Title test 1",
						Description: "Description test",
						Color:       EmbedsColorGreen,
					},
					{
						Title:       "Title test 2",
						Description: "Description test",
						Color:       EmbedsColorRed,
					},
					{
						Title:       "Title test 3",
						Description: "Description test",
						Color:       EmbedsColorYellow,
					},
				},
			},
		},
	)
}
