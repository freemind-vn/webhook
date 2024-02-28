package main

import "freemind.com/webhook/pkg/telegram"

var (
	token  string
	chatId string
)

func Name() string {
	return "github_telegram"
}

func Description() string {
	return "Telegram plugin for GitHub"
}

func Version() string {
	return "0.0.1"
}

func Init(config map[string]string) {
	token = config["token"]
	chatId = config["chat_id"]
}

func Run(payload map[string]any) error {
	return telegram.SendTelegramMessage(token, "chat_id", "message")
}
