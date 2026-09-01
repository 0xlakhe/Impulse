package seller

import (
	"context"

	"github.com/0xlakhe/Impluse/internal/database"
)

type Repository struct{
	db *database.Database
}

func NewRepository(db *database.Database) *Repository{
	return &Repository{db: db}
}

func (r *Repository) FindByID(ctx context.Context,sellerID string)(*Seller,error){
	query:=`
		SELECT
			id,
			name,
			persona,
			system_prompt,
			avatar_url
		FROM sellers
		WHERE id = $1
	`
	var seller Seller

	err:=r.db.Pool.QueryRow(ctx,query,sellerID).Scan(&seller.ID,&seller.Name,&seller.Persona,&seller.SystemPrompt,&seller.AvatarURL)
	
	if err!=nil{
		return nil,err
	}
	
	return &seller,nil
}