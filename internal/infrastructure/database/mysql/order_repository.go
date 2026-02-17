package mysql

import (
	"database/sql"
	"ecommerce-api/internal/domain/entity"
	"fmt"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {

	query := `
		INSERT INTO orders (id, item_id, quantity, total)
		VALUES (?, ?, ?, ?)
	`

	_, err := r.DB.Exec(query, order.ID, order.ItemID, order.Quantity, order.Total)
	if err != nil {
		return err
	}

	return r.DB.QueryRow(
		`SELECT created_at FROM orders WHERE id = ?`,
		order.ID,
	).Scan(&order.CreatedAt)
}

func (r *OrderRepository) FindAll(page, limit int) ([]*entity.Order, error) {
	offset := (page - 1) * limit
	query := `SELECT id, item_id, quantity, total, created_at FROM orders LIMIT ? OFFSET ?`
	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error querying payments: %w", err)
	}
	defer rows.Close()

	orders := make([]*entity.Order, 0)
	for rows.Next() {
		order := &entity.Order{}
		if err := rows.Scan(
			&order.ID,
			&order.ItemID,
			&order.Quantity,
			&order.Total,
			&order.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("error scanning payment row: %w", err)
		}
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", err)
	}

	return orders, nil
}
