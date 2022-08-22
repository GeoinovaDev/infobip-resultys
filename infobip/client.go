package infobip

import (
	"github.com/GeoinovaDev/infobip-resultys/message"
	"github.com/GeoinovaDev/infobip-resultys/webhook"
)

// Client struct
type Client struct {
	APIKey  string
	Webhook *webhook.Server
}

// New cria client
func New(apikey string) *Client {
	return &Client{APIKey: apikey, Webhook: webhook.New(":36465")}
}

// Log ...
func (client *Client) Log(messageID string) ([]message.Message, string) {
	text, err := client.CreateRequest("/tts/3/logs?messageId=" + messageID).Get()
	return client.ProcessResultsResponse(text, err), text
}

// Report ...
func (client *Client) Report(messageID string) ([]message.Message, string) {
	text, err := client.CreateRequest("/tts/3/reports?messageId=" + messageID).Get()
	return client.ProcessResultsResponse(text, err), text
}
