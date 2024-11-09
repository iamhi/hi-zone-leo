package ollamahandler

type MessageResponse struct {
	Role      string          `json:"role"`
	Content   string          `json:"content"`
	ToolCalls []ToolsResponse `json:"tool_calls"`
}
