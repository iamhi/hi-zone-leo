package ollamahandler

import (
	ollamahandler "github.com/iamhi/leo/internal/ollamahandler/requests"
)

func getAllTools() []ollamahandler.ToolRequest {
	return []ollamahandler.ToolRequest{
		{
			Type: "function",
			Function: ollamahandler.FunctionRequest{
				Name:        "get_current_weather",
				Description: "Get the current weather for location",
				Parameters: ollamahandler.ParameterRequest{
					Type: "object",
					Properties: map[string]any{
						"location": &ollamahandler.PropertyBasicRequest{
							Type:        "string",
							Description: "The location to get the weather for, e.g. San Francisco",
						},
						"format": ollamahandler.PropertyEnumRequest{
							Type:        "string",
							Description: "The format to  return the weather in, e.g. 'celsius' or 'fahrenheit'",
							Enum:        []string{"celsius", "fahrenheit"},
						},
					},
					Required: []string{"location"},
				},
			},
		},
		{
			Type: "function",
			Function: ollamahandler.FunctionRequest{
				Name:        "write_note",
				Description: "Write a note",
				Parameters: ollamahandler.ParameterRequest{
					Type: "object",
					Properties: map[string]any{
						"title": &ollamahandler.PropertyBasicRequest{
							Type:        "string",
							Description: "The title for the note",
						},
						"content": ollamahandler.PropertyBasicRequest{
							Type:        "string",
							Description: "The content of the note",
						},
					},
					Required: []string{"content"},
				},
			},
		},
	}
}
