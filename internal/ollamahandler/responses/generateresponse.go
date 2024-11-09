package ollamahandler

type GenerateResponse struct {
	Model              string  `json:"model"`
	CreatedAt          string  `json:"create_at"`
	Response           string  `json:"response"`
	Done               bool    `json:"done"`
	Context            []int32 `json:"context"`
	TotalDuration      int64   `json:"total_duration"`
	LoadDuration       int64   `json:"load_duration"`
	PromptEvalCount    int16   `json:"prompt_eval_count"`
	PromptEvalDuration int64   `json:"prompt_eval_duration"`
	EvalCount          int32   `json:"eval_count"`
	EvalDuration       int64   `json:"eval_duration"`
}
