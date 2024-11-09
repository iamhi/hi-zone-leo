package ollamahandler

import ollamahandler "github.com/iamhi/leo/internal/ollamahandler/requests"

func createChatWithToolsRequest(content string, history []ollamahandler.MessageRequest) ollamahandler.ChatRequest {
	return ollamahandler.ChatRequest{
		Model: "qwen2.5-coder:1.5b",
		Messages: append(history, ollamahandler.MessageRequest{
			Role:    "user",
			Content: content,
		}),
		Stream: false,
		Tools:  getAllTools(),
	}
}

func createChatRequest(content string) ollamahandler.ChatRequest {
	return ollamahandler.ChatRequest{
		Model: "qwen2.5-coder:1.5b",
		Messages: []ollamahandler.MessageRequest{
			{
				Role:    "user",
				Content: content,
			},
		},
		Stream: false,
	}
}

func createGenerateRequest(prompt string) ollamahandler.GenerateRequest {
	return ollamahandler.GenerateRequest{
		Model:  "qwen2.5-coder:1.5b",
		Prompt: prompt,
	}
}
