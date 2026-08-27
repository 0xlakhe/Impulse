package conversation

import "time"

type CreateConversationRequest struct{
	ProductID *string `json:"product_id,omitempty"`
}

type CreateConversationResponse struct{
	ConversationID string `json:"conversation_id"`
}

type SendMessageRequest struct{
	Content string `json:"content"`
}

type MessageResponse struct{
	ID string `json:"id"`
	ConversationID string `json:"conversation_id"`
	Role string `json:"role"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
