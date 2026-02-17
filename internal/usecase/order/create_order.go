package order

import (
	"ecommerce-api/internal/domain/entity"
	"ecommerce-api/internal/domain/repository"

	"github.com/google/uuid"
)

type CreateOrder struct {
	Repo repository.OrderRepository
}

func NewCreateOrderUseCase(repo repository.OrderRepository) *CreateOrder {
	return &CreateOrder{
		Repo: repo,
	}
}

func (uc *CreateOrder) Execute(itemID string, quantity int, price float64) (*entity.Order, error) {

	order := &entity.Order{
		ID:       uuid.NewString(),
		ItemID:   itemID,
		Quantity: quantity,
		Total:    float64(quantity) * price,
	}

	err := uc.Repo.Save(order)
	if err != nil {
		return nil, err
	}

	return order, nil
}
