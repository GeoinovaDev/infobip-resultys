package ivr

import (
	"git.resultys.com.br/lib/lower/convert/decode"
	"git.resultys.com.br/lib/lower/net/request"
	"git.resultys.com.br/sdk/infobip-golang/response"
)

// IVR struct
type IVR struct {
	APIKey     string
	BulkID     string
	ScenarioID string
	To         string
	From       string
	Webhook    string
	Parameters map[string]string
}

type ivrJSONMessage struct {
	ScenarioID   string              `json:"scenarioId"`
	From         string              `json:"from"`
	Destinations []map[string]string `json:"destinations"`
	// NotifyURL         string              `json:"notifyUrl"`
	// NotifyContentType string              `json:"notifyContentType"`
	ValidityPeriod int               `json:"validityPeriod"`
	Retry          map[string]int    `json:"retry"`
	Parameters     map[string]string `json:"parameters"`
}

type ivrJSON struct {
	BuldID   string         `json:"bulkId"`
	Messages ivrJSONMessage `json:"messages"`
}

// New cria um IVR
func New(apikey string) *IVR {
	return &IVR{APIKey: apikey, Parameters: make(map[string]string)}
}

// Call realiza a chamada
func (ivr *IVR) Call() *response.Response {
	obj := ivrJSON{}
	obj.BuldID = ivr.BulkID
	obj.Messages.ScenarioID = ivr.ScenarioID
	obj.Messages.From = ivr.From

	obj.Messages.Destinations = make([]map[string]string, 1)
	obj.Messages.Destinations[0] = make(map[string]string)
	obj.Messages.Destinations[0]["to"] = ivr.To
	// obj.Messages.NotifyURL = ivr.Webhook
	// obj.Messages.NotifyContentType = "application/json"
	obj.Messages.ValidityPeriod = 720

	obj.Messages.Retry = make(map[string]int)
	// obj.Messages.Retry["minPeriod"] = 1
	// obj.Messages.Retry["maxPeriod"] = 5
	obj.Messages.Retry["maxCount"] = 0

	if ivr.Parameters != nil {
		obj.Messages.Parameters = make(map[string]string)
		for k, v := range ivr.Parameters {
			obj.Messages.Parameters[k] = v
		}
	}

	text, err := ivr.createRequest("https://api.infobip.com/voice/ivr/1/messages").PostJSON(&obj)

	return processResponse(text, err)
}

// Log busca o log da ligação
func (ivr *IVR) Log(messageID string) {
	text, err := ivr.createRequest("http://api.infobip.com/tts/3/logs?messageId=" + messageID).Get()
	response := processResponse(text, err)
	response.
}

func processResponse(text string, err error) *response.Response {
	response := response.New()
	if err != nil {
		response.Messages[0].Error.GroupID = 100
		response.Messages[0].Error.Description = err.Error()
	} else {
		decode.JSON(text, &response)
	}

	return response
}

func (ivr *IVR) createRequest(url string) *request.CURL {
	request := request.New(url)
	request.AddHeader("accept", "application/json")
	request.AddHeader("authorization", "Basic "+ivr.APIKey)
	request.AddHeader("content-type", "application/json")
	request.SetTimeout(30)

	return request
}
