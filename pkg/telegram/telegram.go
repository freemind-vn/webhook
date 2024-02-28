package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
)

var specialChars = []string{
	"\\",
	"_",
	"*",
	"[",
	"]",
	"(",
	")",
	"~",
	"`",
	">",
	"<",
	"&",
	"#",
	"+",
	"-",
	"=",
	"|",
	"{",
	"}",
	".",
	"!",
}

func SendTelegramMessage(token string, chatId string, message string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
	data := map[string]string{
		"chat_id":              chatId,
		"text":                 escapeTelegramMessage(message),
		"parse_mode":           "MarkdownV2",
		"disable_notification": "true",
	}

	buf, err := json.Marshal(data)
	if err != nil {
		return err
	}

	res, err := http.Post(url, "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("telegram error: %s", res.Status)
	}

	defer res.Body.Close()

	// if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
	// 	return err
	// }

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	log.Info().Msg(string(body))
	return nil
}

func escapeTelegramMessage(message string) string {
	for _, v := range specialChars {
		message = strings.ReplaceAll(message, v, fmt.Sprintf("\\%s", v))
	}
	return message
}
