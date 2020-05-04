package call

import (
	"errors"
	"time"

	"git.resultys.com.br/lib/lower/exec"
	"git.resultys.com.br/sdk/infobip-golang/infobip"
	"git.resultys.com.br/sdk/infobip-golang/message"
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
func (client *Client) Wait() (msg message.Message, err error) {
	exec.Try(func() {
		client.waiting = true
		count := 0

		client.infobip.Webhook.AddHook(client.MessageID).Ok(func(r interface{}) {
			msg = r.(message.Message)
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
			if count == 120 {
				err = errors.New("timeout de 5 min: MessageID = " + client.MessageID)
				break
			}
		}

		if !client.infobip.Webhook.ExistHook(client.MessageID) {
			client.infobip.Webhook.RemoveHook(client.MessageID)
		}

	}).Catch(func(msg string) {
		err = errors.New(msg)
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
