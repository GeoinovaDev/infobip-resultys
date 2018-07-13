package call

import (
	"errors"
	"time"

	"git.resultys.com.br/lib/lower/exec"
	"git.resultys.com.br/sdk/infobip-golang/infobip"
	"git.resultys.com.br/sdk/infobip-golang/message"
	"git.resultys.com.br/sdk/infobip-golang/response"
)

// Client struct
type Client struct {
	MessageID   string
	infobip     *infobip.Client
	LastMessage string

	waiting bool
}

// New ...
func New(messageID string, infobip *infobip.Client) *Client {
	return &Client{MessageID: messageID, infobip: infobip}
}

// Wait ...
func (client *Client) Wait() (message message.Message, err error) {
	exec.Try(func() {
		client.waiting = true
		count := 0

		client.infobip.Webhook.AddHook(client.MessageID).Ok(func(r interface{}) {
			result := r.(response.ResultsResponse)
			message = result.Messages[0]
			err = nil
			client.waiting = false
		})

		for {
			time.Sleep(1 * time.Second)
			if !client.waiting {
				break
			}
			count++

			// espera até  5 min
			if count == 60*5 {
				err = errors.New("timeout de 5 min")
				break
			}
		}

	}).Catch(func(message string) {
		err = errors.New(message)
	})

	return
}

// WaitPolling ...
func (client *Client) WaitPolling() message.Message {
	for {
		logMessages, _ := client.infobip.Log(client.MessageID)
		if len(logMessages) == 0 {
			time.Sleep(1 * time.Second)
			continue
		}

		message := logMessages[0]
		if message.Error.ID == 0 {
			time.Sleep(1 * time.Second)
			continue
		}

		reportMessages, json := client.infobip.Report(client.MessageID)
		if len(reportMessages) == 0 {
			return message
		}

		client.LastMessage = json

		return reportMessages[0]
	}
}
