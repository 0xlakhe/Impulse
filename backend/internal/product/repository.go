package product

import (
	"context"

	"github.com/0xlakhe/Impluse/internal/database"
)

type Repository struct {
	db *database.Database
}

func NewRepository(db *database.Database) *Repository {
	return &Repository{db: db}
}

func (r *Repository) List(ctx context.Context) ([]ProductResponse, error) {
	query := `
		SELECT
 		   	p.id,
		    p.name,
		    p.description,
		    p.price,
		    p.image_url,
		    p.category,
		    s.name
		FROM products p
		JOIN sellers s
		    ON s.id = p.seller_id
		WHERE p.is_active = TRUE
		ORDER BY p.created_at DESC;
		`

	rows, err := r.db.Pool.Query(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []ProductResponse
	for rows.Next() {
		var p ProductResponse
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.Price,
			&p.ImageURL,
			&p.Category,
			&p.SellerName,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *Repository) FindByID(ctx context.Context, productID string) (*Product, error) {
	query := `
		SELECT
			id,
			seller_id,
			name,
			description,
			price,
			image_url,
			category,
			is_active,
			created_at,
			updated_at
		FROM products
		WHERE id=$1
	`
	var product Product

	err := r.db.Pool.QueryRow(ctx, query, productID).Scan(
		&product.ID,
		&product.SellerID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.ImageURL,
		&product.Category,
		&product.IsActive,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
