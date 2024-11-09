package ollamahandler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/iamhi/leo/internal/errors"
	ollamahandler_requests "github.com/iamhi/leo/internal/ollamahandler/requests"
	ollamahandler_responses "github.com/iamhi/leo/internal/ollamahandler/responses"
)

const application_json_media_type = "application/json"

func SendChatMessage(message string,
	history []ollamahandler_requests.MessageRequest) (ollamahandler_responses.ChatResponse, errors.OllamaHandlerError) {
	request_body, err := json.Marshal(createChatWithToolsRequest(message, history))

	if err != nil {
		fmt.Printf("Unable to marshal request: %s", err)

		return ollamahandler_responses.ChatResponse{}, errors.OllamaGenericError{}
	}

	response, err := http.Post(
		"http://wildberry.local:11434/api/chat",
		application_json_media_type,
		bytes.NewBuffer(request_body))

	if err != nil {
		fmt.Printf("There was an error getting a response %s", err)

		return ollamahandler_responses.ChatResponse{}, errors.OllamaGenericError{}
	}

	defer response.Body.Close()

	var response_body ollamahandler_responses.ChatResponse

	err = json.NewDecoder(response.Body).Decode(&response_body)

	if err != nil {
		fmt.Printf("There was an error decoding the response %s", err)

		return ollamahandler_responses.ChatResponse{}, errors.OllamaGenericError{}
	}

	return response_body, nil
}

// This is mostly used for testing
func SendGenerateMessage(message string) ollamahandler_responses.GenerateResponse {
	request_body, err := json.Marshal(createGenerateRequest(message))

	if err != nil {
		fmt.Printf("Unable to marshal request: %s", err)

		return ollamahandler_responses.GenerateResponse{}
	}

	response, err := http.Post(
		"http://wildberry.local:11434/api/generate",
		application_json_media_type,
		bytes.NewBuffer(request_body))

	if err != nil {
		fmt.Printf("There was an error getting a response %s", err)

		return ollamahandler_responses.GenerateResponse{}
	}

	defer response.Body.Close()

	var response_body ollamahandler_responses.GenerateResponse

	err = json.NewDecoder(response.Body).Decode(&response_body)

	if err != nil {
		fmt.Printf("There was an error decoding the response %s", err)

		return ollamahandler_responses.GenerateResponse{}
	}

	return response_body
}
