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
	Parameters     map[string]string   `json:"parameters"`

	NotifyURL         string `json:"notifyUrl"`
	NotifyContentType string `json:"notifyContentType"`
}

type ivrJSON struct {
	BuldID   string         `json:"bulkId"`
	Messages ivrJSONMessage `json:"messages"`
}

// New ...
func New(client *infobip.Client) *Client {
	return &Client{infobip: client}
}

// Call ...
func (client *Client) Call(ivr *IVR) (*call.Client, error) {
	obj := ivrJSON{}
	obj.BuldID = ivr.BulkID
	obj.Messages.ScenarioID = ivr.ScenarioID
	obj.Messages.From = ivr.From

	obj.Messages.Destinations = make([]map[string]string, 1)
	obj.Messages.Destinations[0] = make(map[string]string)
	obj.Messages.Destinations[0]["to"] = ivr.To

	obj.Messages.NotifyURL = ivr.Webhook
	obj.Messages.NotifyContentType = "application/json"
	obj.Messages.ValidityPeriod = 720

	//obj.Messages.Retry = make(map[string]int)
	//obj.Messages.Retry["maxCount"] = 0

	if ivr.Parameters != nil {
		obj.Messages.Parameters = make(map[string]string)
		for k, v := range ivr.Parameters {
			obj.Messages.Parameters[k] = v
		}
	}

	text, err := client.infobip.CreateRequest("/voice/ivr/1/messages").PostJSON(&obj)
	messages := client.infobip.ProcessCallResponse(text, err)
	if len(messages) == 0 {
		return nil, fmt.Errorf(text)
	}

	return call.New(messages[0].MessageID, client.infobip), nil
}
