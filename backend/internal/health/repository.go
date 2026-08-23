package health

import (
	"context"

	"github.com/0xlakhe/Impluse/internal/database"
)

type Repository struct{
	db *database.Database
}

func NewRepository(db *database.Database) *Repository{
	return &Repository{
		db: db,
	}
}

func (r *Repository) Ping(ctx context.Context) error{
	var result int
	err:=r.db.Pool.QueryRow(
		ctx,
		"SELECT 1",
	).Scan(&result)
	return err
}