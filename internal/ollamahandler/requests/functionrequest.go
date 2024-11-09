package ollamahandler

type FunctionRequest struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Parameters  ParameterRequest `json:"parameters"`
}
