package action

// Collect struct
type Collect struct {
	Variable string         `json:"collectInto"`
	Options  map[string]int `json:"options"`
}
