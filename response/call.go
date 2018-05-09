package response

import "git.resultys.com.br/sdk/infobip-golang/message"

// CallResponse struct
type CallResponse struct {
	BulkID   string            `json:"bulkId"`
	Messages []message.Message `json:"messages"`
}
