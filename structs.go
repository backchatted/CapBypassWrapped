package capbypass

type CapBypass struct {
	apiKey string
}

type CapBypassResponse struct {
	ApiKey           string      `json:"apiKey,omitempty"`
	Credits          float32     `json:"credits,omitempty"`
	TotalSolves      int         `json:"totalSolves,omitempty"`
	Stats            interface{} `json:"stats,omitempty"`
	ErrorId          int         `json:"errorId,omitempty"`
	Status           string      `json:"status,omitempty"`
	Solution         string      `json:"solution,omitempty"`
	TaskId           string      `json:"taskId,omitempty"`
	Errordescription string      `json:"errorDescription,omitempty"`
	ErrorMessage     string      `json:"errorMessage,omitempty"`
}

type CapBypassPayload struct {
	ClientKey string `json:"clientKey,omitempty"`
	Task      struct {
		Type             string          `json:"type,omitempty"`
		WebsiteURL       string          `json:"websiteURL,omitempty"`
		WebsitePublicKey string          `json:"websitePublicKey,omitempty"`
		WebsiteSubdomain string          `json:"websiteSubdomain,omitempty"`
		Proxy            string          `json:"proxy,omitempty"`
		Data             string          `json:"data,omitempty"`
		Headers          *map[string]any `json:"headers,omitempty"`
	} `json:"task,omitempty"`
	TaskId string `json:"taskId,omitempty"`
}
