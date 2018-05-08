package infobip

import (
	"git.resultys.com.br/lib/lower/convert/decode"

	"git.resultys.com.br/lib/lower/net/request"
	"git.resultys.com.br/sdk/infobip-golang/message"
)

type resultsResponse struct {
	Messages []message.Message `json:"results"`
}

type callResponse struct {
	BulkID   string            `json:"bulkId"`
	Messages []message.Message `json:"messages"`
}

// CreateRequest ...
func (client *Client) CreateRequest(url string) *request.CURL {
	request := request.New("https://api.infobip.com" + url)
	request.AddHeader("accept", "application/json")
	request.AddHeader("authorization", "Basic "+client.APIKey)
	request.AddHeader("content-type", "application/json")
	request.SetTimeout(30)

	return request
}

// ProcessCallResponse ...
func (client *Client) ProcessCallResponse(text string, err error) []message.Message {
	results := callResponse{Messages: make([]message.Message, 1)}

	if err != nil {
		results.Messages[0].Error.ID = 100
		results.Messages[0].Error.Description = err.Error()
	} else {
		decode.JSON(text, &results)
	}

	return results.Messages
}

// ProcessResultsResponse ...
func (client *Client) ProcessResultsResponse(text string, err error) []message.Message {
	results := resultsResponse{Messages: make([]message.Message, 1)}

	if err != nil {
		results.Messages[0].Error.ID = 100
		results.Messages[0].Error.Description = err.Error()
	} else {
		decode.JSON(text, &results)
	}

	return results.Messages
}
