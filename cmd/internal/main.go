package main

import (
	"encoding/json"
	"fmt"

	"github.com/iamhi/leo/internal/ollamahandler"
	ollama_requests "github.com/iamhi/leo/internal/ollamahandler/requests"
)

func main() {
	fmt.Println("Running examples")

	// exampleChatWithTools()
	exampleGenerate()
}

func exampleChatWithTools() {
	// Writing note
	// chat_response := ollamahandler.SendChatMessage("Write a note with content: \"What is the weather?\" and Title: Weather talk")
	// Getting the weather
	chat_response, _ := ollamahandler.SendChatMessage("Write about your favorite pokemon and add it as note", []ollama_requests.MessageRequest{})
	response_json, err := json.Marshal(chat_response)

	if err == nil {
		fmt.Println(string(response_json))
	} else {
		fmt.Printf("Error while marshaling the request %s", err)
	}
}

func exampleGenerate() {
	chat_response := ollamahandler.SendGenerateMessage("Write about your favorite pokemon and add it as note")
	response_json, err := json.Marshal(chat_response)

	if err == nil {
		fmt.Println(string(response_json))
	} else {
		fmt.Printf("Error while marshaling the request %s", err)
	}
}
