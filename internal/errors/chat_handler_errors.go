package errors

const CHAT_EXISTS_CODE = "chat-exists"
const CHAT_NOT_FOUND_CODE = "chat-not-found"
const CHAT_OLLAMA_ERROR_CODE = "chat-ollama-error"

type ChatHandlerError interface {
	Error() string
	GetCode() string
}

type ChatExistsError struct {
	Username string
}

func (e ChatExistsError) Error() string {
	return "Chat already exists"
}

func (e ChatExistsError) GetCode() string {
	return CHAT_EXISTS_CODE
}

type ChatNotFound struct {
	OwnerUuid string
}

func (e ChatNotFound) Error() string {
	return "Chat not found"
}

func (e ChatNotFound) GetCode() string {
	return CHAT_NOT_FOUND_CODE
}

type ChatOllamaError struct{}

func (e ChatOllamaError) Error() string {
	return "Error occurred while contacting Ollama"
}

func (e ChatOllamaError) GetCode() string {
	return CHAT_OLLAMA_ERROR_CODE
}
