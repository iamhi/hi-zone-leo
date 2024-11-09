package ollamahandler

type PropertyRequest interface {
	PropertyBasicRequest | PropertyEnumRequest
}

type PropertyBasicRequest struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

type PropertyEnumRequest struct {
	Type        string   `json:"type"`
	Description string   `json:"description"`
	Enum        []string `json:"enum"`
}
