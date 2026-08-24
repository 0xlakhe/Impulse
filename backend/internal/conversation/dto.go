package conversation

type CreateConversationRequest struct{
	ProductID *string `json:"product_id,omitempty"`
}

type CreateConversationResponse struct{
	ConversationID string `json:"conversation_id"`
}