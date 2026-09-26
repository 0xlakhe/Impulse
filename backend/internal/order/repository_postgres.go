package order

import (
	"context"
	"fmt"
	"github.com/0xlakhe/Impluse/internal/database"
)

type repository struct {
	db *database.Database
}

func NewRepository(db *database.Database) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, order *Order, items []OrderItem) error {
	tx, err := r.db.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin order transaction: %w", err)
	}

	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx,
		`INSERT INTO orders (
			id,
			user_id,
			total_amount,
			status,
			created_at,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6)`,
		order.ID,
		order.UserID,
		order.TotalAmount,
		order.Status,
		order.CreatedAt,
		order.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("create order: %w", err)
	}

	for _, item := range items {
		_, err = tx.Exec(ctx,
			`INSERT INTO order_items (
				id,
				order_id,
				product_id,
				product_name,
				unit_price,
				quantity,
				subtotal
			)
			VALUES ($1, $2, $3, $4, $5, $6, $7)`,
			item.ID,
			item.OrderID,
			item.ProductID,
			item.ProductName,
			item.UnitPrice,
			item.Quantity,
			item.Subtotal,
		)
		if err != nil {
			return fmt.Errorf("create order item: %w", err)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit order transaction: %w", err)
	}
	return nil
}

func (r *repository) UpdateStatus(ctx context.Context, orderID string, status Status,)error{
	result,err:=r.db.Pool.Exec(ctx,
		`UPDATE orders
		SET
			status = $1,
			updated_at = NOW()
		WHERE id = $2`,
		status,orderID)
	
	if err!=nil{
		return fmt.Errorf("update order status: %w",err)
	}
	if result.RowsAffected()==0{
		return fmt.Errorf("order %s not found",orderID)
	}
	return nil
}