package ollamahandler

type GenerateRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Suffix string `json:"suffix"`
	Stream bool   `json:"stream"`
}
