package go_slack_webhook_client

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

const (
	DefaultTimeout = 5 * time.Second

	TypePlainText = "plain_text"
	TypeMrkdwn    = "mrkdwn"
)

type WCConfig struct {
	Webhook string
	Channel string
	Timeout time.Duration
	Headers map[string]string
}

type WebhookClient struct {
	Config *WCConfig
}

type SubBlock struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
}

type Block struct {
	Type   string     `json:"type,omitempty"`
	Text   SubBlock   `json:"text,omitempty"`
	Fields []SubBlock `json:"fields,omitempty"`
}

type Message struct {
	IconEmoji string  `json:"icon_emoji,omitempty"`
	Channel   string  `json:"channel,omitempty"`
	Blocks    []Block `json:"blocks,omitempty"`
}

func NewWebhookClient(config *WCConfig) *WebhookClient {
	return &WebhookClient{
		Config: config,
	}
}

func (wc *WebhookClient) SendMessage(m *Message) error {
	return wc.send(m)
}

func (wc *WebhookClient) send(m *Message) error {
	body, _ := json.Marshal(m)

	req, err := http.NewRequest(http.MethodPost, wc.Config.Webhook, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	if len(wc.Config.Headers) > 0 {
		for key, value := range wc.Config.Headers {
			req.Header.Add(key, value)
		}
	}

	if wc.Config.Timeout == 0 {
		wc.Config.Timeout = DefaultTimeout
	}

	client := &http.Client{Timeout: wc.Config.Timeout}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return err
	}

	if buf.String() != "ok" {
		return errors.New("Slack notification failed")
	}

	return nil
}
