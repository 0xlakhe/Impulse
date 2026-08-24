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