package ollamahandler

type ChatRequest struct {
	Model    string           `json:"model"`
	Messages []MessageRequest `json:"messages"`
	Stream   bool             `json:"stream"`
	Tools    []ToolRequest    `json:"tools"`
}
