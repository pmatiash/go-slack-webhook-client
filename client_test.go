package go_slack_webhook_client

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewWebhookClient(t *testing.T) {
	config := &WCConfig{
		Webhook:  "test-webhook-url",
		Username: "test-username",
		Channel:  "test-channel",
		Timeout:  2 * time.Second,
	}
	c := NewWebhookClient(config)

	assert.Equal(t, "test-webhook-url", c.Config.Webhook)
	assert.Equal(t, "test-username", c.Config.Username)
	assert.Equal(t, "test-channel", c.Config.Channel)
	assert.Equal(t, 2*time.Second, c.Config.Timeout)
}
