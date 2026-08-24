package conversation

import "time"


type Conversation struct{
	ID string
	UserID string
	SellerID string
	CreatedAt time.Time
	UpdatedAt time.Time
}