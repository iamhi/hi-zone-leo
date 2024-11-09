package chathandler

import (
	"github.com/google/uuid"
	"github.com/iamhi/leo/db/postgres"
	"github.com/iamhi/leo/db/postgres/models"
	"github.com/iamhi/leo/internal/errors"
	"github.com/iamhi/leo/internal/ollamahandler"
	"github.com/iamhi/leo/internal/userhandler"

	ollamahandler_requests "github.com/iamhi/leo/internal/ollamahandler/requests"
)

type ChatDto struct {
	Uuid      string       `json:"uuid"`
	OwnerUuid string       `json:"ownerUuid"`
	Messages  []MessageDto `json:"messages"`
}

type MessageDto struct {
	Uuid     string `json:"uuid"`
	ChatUuid string `json:"chat_uuid"`
	Role     string `json:"role"`
	Content  string `json:"content"`
}

func GetChat(user_details userhandler.UserDetails) (ChatDto, errors.ChatHandlerError) {
	chat, chat_handler_error := findChat(user_details)

	if chat_handler_error != nil {
		chat, chat_handler_error = createChat(user_details)

		if chat_handler_error != nil {
			return ChatDto{}, chat_handler_error
		}
	}

	chat_dto := ChatDto{
		Uuid:      chat.Uuid,
		OwnerUuid: chat.OwnerUuid,
	}

	message_dtos := make([]MessageDto, 0)
	messages_entity := findMessagesForChat(chat_dto)

	for _, message := range messages_entity {
		message_dtos = append(message_dtos, MessageDto{
			Uuid:     message.Uuid,
			ChatUuid: message.ChatUuid,
			Role:     message.Role,
			Content:  message.Content,
		})
	}

	chat_dto.Messages = message_dtos

	return chat_dto, nil
}

func SendMessage(user_details userhandler.UserDetails, content string) (ChatDto, errors.ChatHandlerError) {
	chat_dto, chat_handler_error := GetChat(user_details)

	if chat_handler_error != nil {

		return ChatDto{}, chat_handler_error
	}

	request_history := make([]ollamahandler_requests.MessageRequest, 0)

	for _, message := range chat_dto.Messages {
		request_history = append(request_history, ollamahandler_requests.MessageRequest{
			Role:    message.Role,
			Content: message.Content,
		})
	}

	ollama_response, ollama_handler_error := ollamahandler.SendChatMessage(content, request_history)

	if ollama_handler_error != nil {
		return ChatDto{}, errors.ChatOllamaError{}
	}

	ollama_message_content := ollama_response.Message.Content
	ollama_message_role := "assistant"

	if len(ollama_response.Message.ToolCalls) != 0 {
		// TODO: Handle the tools
		ollama_message_content = ollama_response.Message.ToolCalls[0].Function.Name
		ollama_message_role = "tool"
	}

	user_message := models.Message{
		Uuid:     string(uuid.New().String()),
		ChatUuid: chat_dto.Uuid,
		Role:     "user",
		Content:  content,
	}

	ollama_message := models.Message{
		Uuid:     string(uuid.New().String()),
		ChatUuid: chat_dto.Uuid,
		Role:     ollama_message_role,
		Content:  ollama_message_content,
	}

	// TODO: Handle errors here
	postgres.Db.Create(&user_message)
	postgres.Db.Create(&ollama_message)

	chat_dto.Messages = append(chat_dto.Messages, MessageDto{
		Uuid:     user_message.Uuid,
		ChatUuid: chat_dto.Uuid,
		Role:     user_message.Role,
		Content:  user_message.Content,
	})

	chat_dto.Messages = append(chat_dto.Messages, MessageDto{
		Uuid:     ollama_message.Uuid,
		ChatUuid: chat_dto.Uuid,
		Role:     ollama_message.Role,
		Content:  ollama_message.Content,
	})

	return chat_dto, nil
}

func findChat(user_details userhandler.UserDetails) (models.Chat, errors.ChatHandlerError) {
	var chat models.Chat

	postgres.Db.Model(&models.Chat{}).Where("owner_uuid=?", user_details.Uuid).First(&chat)

	if chat.ID == 0 {
		chat_not_found_error := errors.ChatNotFound{}

		chat_not_found_error.OwnerUuid = user_details.Uuid

		return models.Chat{}, chat_not_found_error
	}

	return chat, nil
}

func findMessagesForChat(chat_dto ChatDto) []models.Message {
	var messages []models.Message

	postgres.Db.Model(&models.Message{}).Where("chat_uuid=?", chat_dto.Uuid).Find(&messages)

	return messages
}

func createChat(user_details userhandler.UserDetails) (models.Chat, errors.ChatHandlerError) {
	_, chat_handler_error := findChat(user_details)

	if chat_handler_error == nil {
		chatExistsError := errors.ChatExistsError{}

		chatExistsError.Username = user_details.Username

		return models.Chat{}, chatExistsError
	}

	new_chat := models.Chat{
		Uuid:      uuid.New().String(),
		OwnerUuid: user_details.Uuid,
	}

	postgres.Db.Create(&new_chat)

	return new_chat, nil
}
