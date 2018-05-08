package action

// Say struct
type Say struct {
	Phrase  string            `json:"say"`
	Options map[string]string `json:"options"`
}
