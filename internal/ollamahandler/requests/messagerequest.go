package ollamahandler

type MessageRequest struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
