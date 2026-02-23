package order

import (
	"ecommerce-api/internal/domain/repository"
)

type UpdateOrderStatusInput struct {
	ID     string
	Status string
}

type UpdateOrderStatus struct {
	Repo repository.OrderRepository
}

func NewUpdateOrderStatusUseCase(repo repository.OrderRepository) *UpdateOrderStatus {
	return &UpdateOrderStatus{
		Repo: repo,
	}
}

func (uc *UpdateOrderStatus) Execute(input UpdateOrderStatusInput) error {
	order, err := uc.Repo.FindByID(input.ID)
	if err != nil {
		return err
	}

	if order == nil {
		return &repository.ErrNotFound{Message: "order not found"}
	}

	order.Status = input.Status

	return uc.Repo.UpdateStatus(order)
}
