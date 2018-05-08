package infobip

import (
	"git.resultys.com.br/sdk/infobip-golang/message"
)

// Client struct
type Client struct {
	APIKey string
}

// New cria client
func New(apikey string) *Client {
	return &Client{APIKey: apikey}
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
