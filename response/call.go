package response

import "github.com/GeoinovaDev/infobip-resultys/message"

// CallResponse struct
type CallResponse struct {
	BulkID   string            `json:"bulkId"`
	Messages []message.Message `json:"messages"`
}
