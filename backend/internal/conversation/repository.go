package conversation

import (
	"context"
	"errors"

	"github.com/0xlakhe/Impluse/internal/database"
	"github.com/jackc/pgx/v5"
)

type Repository struct{
	db *database.Database
}

func NewRepository(db *database.Database) *Repository{
	return &Repository{db: db}
}

func (r *Repository) FindByUserAndSeller(ctx context.Context,userID string, sellerID string)(*Conversation,error){
	query:=`
		SELECT
			id,
			user_id,
			seller_id,
			created_at,
			updated_at
		FROM conversations
		WHERE user_id=$1
		AND seller_id=$2
	`
	var conversation Conversation
	err:=r.db.Pool.QueryRow(ctx,query,userID,sellerID).Scan(&conversation.ID,&conversation.UserID	,&conversation.SellerID,&conversation.CreatedAt,&conversation.UpdatedAt)
	
	if err!=nil{
		return nil,err
	}
	return &conversation,nil

}

func(r *Repository) CreateOrGet(ctx context.Context,userID string,sellerID string,)(*Conversation, error){
	query:=`
	INSERT INTO conversations(
		user_id,
		seller_id
	)
	VALUES(
		$1,
		$2
	)
	ON CONFLICT(
		user_id,
		seller_id
	)
	DO UPDATE SET
		updated_at=conversations.updated_at
	RETURNING
		id,
		user_id,
		seller_id,
		created_at,
		updated_at
	`
	var conversation Conversation
	err:=r.db.Pool.QueryRow(ctx,query,userID,sellerID).Scan(&conversation.ID,&conversation.UserID,&conversation.SellerID,&conversation.CreatedAt,&conversation.UpdatedAt)
	if err!=nil{
		if errors.Is(err,pgx.ErrNoRows){
			return nil,ErrNotFound
		}
		return nil,err
	}
	return &conversation,nil
}

func (r *Repository) CreateMessage(ctx context.Context, conversationID string, role string, content string,)(*Message, error){
	query:=`
		INSERT INTO messages(
			conversation_id,
			role,
			content
		)
		VALUES ($1,$2,$3)
		RETURNING
			id,
			conversation_id,
			role,
			content,
			created_at
		`
	var message Message
	err:=r.db.Pool.QueryRow(ctx,query,conversationID,role,content).Scan(&message.ID,&message.ConversationID,&message.Role,&message.Content,&message.CreatedAt)
	if err!=nil{
		return nil,err
	}
	return &message,nil
}

func(r *Repository) GetMessages(ctx context.Context,conversationID string,)([]Message,error){
	query:=`
		SELECT
			id,
			conversation_id,
			role,
			content,
			created_at
		FROM messages
		WHERE conversation_id=$1
		ORDER BY created_at ASC
	`
	rows,err:=r.db.Pool.Query(
		ctx,query,conversationID,
	)
	if err!=nil{
		return nil,err
	}
	defer rows.Close()
	messages:=make([]Message,0)

	for rows.Next(){
		var message Message
		err:=rows.Scan(
			&message.ID,
			&message.ConversationID,
			&message.Role,
			&message.Content,
			&message.CreatedAt,
		)
		if err!=nil{
			return nil,err
		}
		messages=append(messages, message)		
	}

	if err:=rows.Err();err!=nil{
		return nil,err
	}
	return messages,nil
}

func(r *Repository) BelongsToUser(ctx context.Context, conversationID string, userID string)(bool,error){
	query:=`
		SELECT EXISTS(
			SELECT 1
			FROM conversations
			WHERE id=$1
			AND user_id=$2
		)
	`
	var exists bool
	err:=r.db.Pool.QueryRow(
		ctx,query,conversationID,userID,
	).Scan(&exists)
	if err!=nil{
		return false, err
	}
	return exists,nil
}
