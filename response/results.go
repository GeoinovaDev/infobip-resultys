package response

import "git.resultys.com.br/sdk/infobip-golang/message"

// ResultsResponse struct
type ResultsResponse struct {
	Messages []message.Message `json:"results"`
}
