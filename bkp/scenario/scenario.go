package scenario

import (
	"git.resultys.com.br/lib/lower/convert/decode"
	"git.resultys.com.br/lib/lower/net/request"
	"git.resultys.com.br/sdk/infobip-golang/response"
)

// Scenario struct
type Scenario struct {
	Name    string        `json:"name"`
	Actions []interface{} `json:"script"`
}

// New cria um script
func New() *Scenario {
	return &Scenario{}
}

// AddAction adiciona uma action
func (s *Scenario) AddAction(action interface{}) *Scenario {
	s.Actions = append(s.Actions, action)

	return s
}

// AddHangup adiciona action hangup
func (s *Scenario) AddHangup() *Scenario {
	s.Actions = append(s.Actions, "hangup")

	return s
}

// Insert insere um script
func (s *Scenario) Insert() *response.Response {
	request := request.New("https://api.infobip.com/voice/ivr/1/scenarios")
	request.AddHeader("accept", "application/json")
	request.AddHeader("authorization", "Basic UmVzdWx0eXM6cmVzdWx0eXNAMjUyOQ==")
	request.AddHeader("content-type", "application/json")
	request.SetTimeout(30)
	text, err := request.PostJSON(s)

	resp := &response.Response{}
	if err != nil {
		resp.Messages = make([]response.Message, 1)
		resp.Messages[0].Status.GroupID = 100
		resp.Messages[0].Status.Description = err.Error()
	} else {
		decode.JSON(text, &resp)
	}

	return resp
}
