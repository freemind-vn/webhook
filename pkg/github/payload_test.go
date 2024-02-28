package github

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPayloadPush(t *testing.T) {
	buf, err := os.ReadFile("push.json")
	assert.NoError(t, err)

	var payload Payload
	err = json.Unmarshal(buf, &payload)
	assert.NoError(t, err)
	assert.NotEmpty(t, payload)

	assert.Equal(t, "svelte-admin", payload.Repository.Name)
	assert.Equal(t, "https://github.com/freemind-vn/svelte-admin/compare/e7f1ad5c82ac...4bded6500b83", payload.Compare)
}
