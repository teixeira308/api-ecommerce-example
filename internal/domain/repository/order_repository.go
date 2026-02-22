package repository

import "ecommerce-api/internal/domain/entity"

type OrderRepository interface {
	Save(*entity.Order) error
	FindAll(page, limit int) ([]*entity.Order, error)
	FindByID(id string) (*entity.Order, error)
}
