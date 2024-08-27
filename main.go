package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const botToken = "7439590358:AAFN5dC58gT_2UXlcePTkLrSWmdmVNF4wUc"
const chatID = "291109889"

type WebAppInfo struct {
	URL string `json:"url"`
}

type InlineKeyboardButton struct {
	Text   string     `json:"text"`
	WebApp *WebAppInfo `json:"web_app"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type SendMessageRequest struct {
	ChatID      string               `json:"chat_id"`
	Text        string               `json:"text"`
	ReplyMarkup InlineKeyboardMarkup `json:"reply_markup"`
}

func sendWebAppButton() error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

	webApp := WebAppInfo{URL: "https://calixtemayoraz.gitlab.io/web-interfacer-bot/"}
	button := InlineKeyboardButton{
		Text:   "Open Web App",
		WebApp: &webApp,
	}
	keyboard := InlineKeyboardMarkup{
		InlineKeyboard: [][]InlineKeyboardButton{{button}},
	}
	requestBody := SendMessageRequest{
		ChatID:      chatID,
		Text:        "Click the button to open the Web App",
		ReplyMarkup: keyboard,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send message: %s", resp.Status)
	}

	return nil
}

func main() {
	if err := sendWebAppButton(); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Web app button sent successfully.")
	}
}
