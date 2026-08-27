package conversation

func toMessageResponse(
	message Message,
)MessageResponse{
	return MessageResponse{
		ID: message.ID,
		ConversationID: message.ConversationID,
		Role: message.Role,
		Content: message.Content,
		CreatedAt: message.CreatedAt,
	}
}