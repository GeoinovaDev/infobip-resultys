package response

import "github.com/GeoinovaDev/infobip-resultys/message"

// ResultsResponse struct
type ResultsResponse struct {
	Messages []message.Message `json:"results"`
}
