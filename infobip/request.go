package infobip

import (
	"github.com/GeoinovaDev/lower-resultys/convert/decode"
	"github.com/GeoinovaDev/lower-resultys/net/request"
	"github.com/GeoinovaDev/infobip-resultys/message"
	"github.com/GeoinovaDev/infobip-resultys/response"
)

// URL Infobip
var URL = "https://api.infobip.com"

// CreateRequest ...
func (client *Client) CreateRequest(url string) *request.CURL {
	request := request.New(URL + url)
	request.AddHeader("accept", "application/json")
	request.AddHeader("authorization", "Basic "+client.APIKey)
	request.AddHeader("content-type", "application/json")
	request.SetTimeout(30)

	return request
}

// ProcessCallResponse ...
func (client *Client) ProcessCallResponse(text string, err error) []message.Message {
	results := response.CallResponse{Messages: make([]message.Message, 1)}

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
	results := response.ResultsResponse{Messages: make([]message.Message, 1)}

	if err != nil {
		results.Messages[0].Error.ID = 100
		results.Messages[0].Error.Description = err.Error()
	} else {
		decode.JSON(text, &results)
	}

	return results.Messages
}
