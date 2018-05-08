package message

import (
	"strconv"
	"strings"
)

// Price struct
type Price struct {
	PricePerSecond float64 `json:"pricePerSecond"`
	Currency       string  `json:"currency"`
}

// Status struct
type Status struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	GroupID     int    `json:"groupId"`
	GroupName   string `json:"groupName"`
}

// Message struct
type Message struct {
	Status          Status `json:"status"`
	Error           Status `json:"error"`
	BulkID          string `json:"bulkId"`
	MessageID       string `json:"messageId"`
	To              string `json:"to"`
	From            string `json:"from"`
	SentAt          string `json:"sentAt"`
	DoneAt          string `json:"doneAt"`
	Duration        int    `json:"duration"`
	MccMnc          string `json:"mccMnc"`
	Price           Price  `json:"price"`
	RecordedFileURL string `json:"recordedFileUrl"`
	DtmfCodes       string `json:"dtmfCodes,omitempty"`
}

// KeyPressNumber ...
func (m *Message) KeyPressNumber() int {
	if len(m.DtmfCodes) == 0 {
		return -1
	}

	n, e := strconv.Atoi(strings.Replace(m.DtmfCodes, ",", "", -1))
	if e != nil {
		return -1
	}

	return n
}
