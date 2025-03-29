package bot

import (
	"fmt"
	"net/http"
	"net/url"
)

func SendTelegramMessage(chatID, message, botToken string) {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

	data := url.Values{}
	data.Set("chat_id", chatID)
	data.Set("text", message)

	_, err := http.PostForm(apiURL, data)
	if err != nil {
		fmt.Println("Error sending message:", err)
	} else {
		fmt.Println("Telegram message sent!")
	}
}
