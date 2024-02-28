package telegram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendTelegramMessage(t *testing.T) {
	err := SendTelegramMessage("token", "chat_id", "message")
	assert.Error(t, err)
}
