package errors

const OLLAMA_GENERIC_ERROR_CODE = "ollama-generic-error"

type OllamaHandlerError interface {
	GetCode() string
}

type OllamaGenericError struct {
}

func (e OllamaGenericError) Error() string {
	return "Ollama generic error"
}

func (e OllamaGenericError) GetCode() string {
	return OLLAMA_GENERIC_ERROR_CODE
}
