package response

// Response body response
type Response struct {
	BulkID   string    `json:"bulkId,omitempty"`
	Messages []Message `json:"messages"`
	ID       string    `json:"id"`
}

// Message ...
type Message struct {
	ID     string `json:"messageId"`
	Status Status `json:"status"`
	Error  Status `json:"error"`
}

// Status ...
type Status struct {
	ID          int    `json:"id"`
	GroupID     int    `json:"groupId"`
	GroupName   string `json:"groupName"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// New cria a resposta
func New() *Response {
	return &Response{Messages: make([]Message, 0)}
}
