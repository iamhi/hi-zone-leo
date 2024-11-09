package ollamahandler

type FunctionResponse struct {
	Name      string         `json:"name"`
	Arguments map[string]any `json:"arguments"`
}
