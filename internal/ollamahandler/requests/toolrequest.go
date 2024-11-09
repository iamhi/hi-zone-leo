package ollamahandler

type ToolRequest struct {
	Type     string          `json:"type"`
	Function FunctionRequest `json:"function"`
}
