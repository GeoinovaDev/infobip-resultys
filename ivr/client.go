package ivr

import (
	"fmt"

	"git.resultys.com.br/sdk/infobip-golang/call"
	"git.resultys.com.br/sdk/infobip-golang/infobip"
)

// Client struct
type Client struct {
	infobip *infobip.Client
}

// IVR struct
type IVR struct {
	BulkID     string
	ScenarioID string
	To         string
	From       string
	Webhook    string
	Parameters map[string]string
}

type ivrJSONMessage struct {
	ScenarioID     string              `json:"scenarioId"`
	From           string              `json:"from"`
	Destinations   []map[string]string `json:"destinations"`
	ValidityPeriod int                 `json:"validityPeriod"`
	//Retry          map[string]int      `json:"retry"`
	Parameters map[string]string `json:"parameters"`

	NotifyURL         string `json:"notifyUrl"`
	NotifyContentType string `json:"notifyContentType"`
}

type ivrJSON struct {
	BuldID   string           `json:"bulkId"`
	Messages []ivrJSONMessage `json:"messages"`
}

// New ...
func New(client *infobip.Client) *Client {
	return &Client{infobip: client}
}

// Call ...
func (client *Client) Call(ivr *IVR) (*call.Client, error) {
	obj := ivrJSON{}
	obj.BuldID = ivr.BulkID
	message := ivrJSONMessage{}
	message.ScenarioID = ivr.ScenarioID
	message.From = ivr.From

	message.Destinations = make([]map[string]string, 1)
	message.Destinations[0] = make(map[string]string)
	message.Destinations[0]["to"] = ivr.To

	message.NotifyURL = ivr.Webhook
	message.NotifyContentType = "application/json"
	message.ValidityPeriod = 720

	//message.Retry = make(map[string]int)
	//message.Retry["maxCount"] = 0

	if ivr.Parameters != nil {
		message.Parameters = make(map[string]string)
		for k, v := range ivr.Parameters {
			message.Parameters[k] = v
		}
	}

	obj.Messages = []ivrJSONMessage{message}

	text, err := client.infobip.CreateRequest("/voice/ivr/1/messages").PostJSON(&obj)
	messages := client.infobip.ProcessCallResponse(text, err)
	if len(messages) == 0 {
		return nil, fmt.Errorf(text)
	}

	return call.New(messages[0].MessageID, client.infobip), nil
}
