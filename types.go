package capbypass

type CapBypass struct {
	apiKey string
}

type CapBypassResponse struct {
	ApiKey      string      `json:"apiKey,omitempty"`
	Credits     float32     `json:"credits,omitempty"`
	TotalSolves int         `json:"totalSolves,omitempty"`
	Stats       interface{} `json:"stats,omitempty"`
	ErrorId     int         `json:"errorId,omitempty"`
	Status      string      `json:"status,omitempty"`
	Solution    string      `json:"solution,omitempty"`
	TaskId      string      `json:"taskId,omitempty"`
}

type CapBypassPayload struct {
	ClientKey string          `json:"clientKey,omitempty"`
	Task      *map[string]any `json:"task,omitempty"`
	TaskId    string          `json:"taskId,omitempty"`
}
