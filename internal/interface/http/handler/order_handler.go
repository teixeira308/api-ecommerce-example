package handler

import (
	"ecommerce-api/internal/interface/dto"
	"ecommerce-api/internal/usecase/order"
	"encoding/json"
	"errors"
	"net/http"
)

// ErrorResponse represents a standardized JSON error response
type OrderErrorResponse struct {
	Message string `json:"message"`
}

// orderRespondWithError sends a JSON error response
func orderRespondWithError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(OrderErrorResponse{Message: message})
}

type OrderHandler struct {
	CreateOrder  *order.CreateOrder
	GetAllOrders *order.GetAllOrders
	GetOrder     *order.GetOrder
}

func NewOrderHandler(
	createOrder *order.CreateOrder,
	getAllOrders *order.GetAllOrders,
	getOrder *order.GetOrder,
) *OrderHandler {
	return &OrderHandler{
		CreateOrder:  createOrder,
		GetAllOrders: getAllOrders,
		GetOrder:     getOrder,
	}
}

func (h *OrderHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// For now, use default page and limit. In a real app, these would come from query parameters.
	input := order.GetAllOrdersInput{
		Page:  1,
		Limit: 10,
	}

	output, err := h.GetAllOrders.Execute(input)
	if err != nil {
		orderRespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := make([]dto.GetOrderResponse, len(output.Orders))
	for i, order := range output.Orders {
		response[i] = dto.GetOrderResponse{
			ID:        order.ID,
			ItemID:    order.ItemID,
			Quantity:  order.Quantity,
			Total:     order.Total,
			CreatedAt: order.CreatedAt,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateOrderRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		orderRespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	order, err := h.CreateOrder.Execute(input.ItemID, input.Quantity, input.Total)
	if err != nil {
		orderRespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	response := dto.CreateOrderResponse{
		ID:        order.ID,
		ItemID:    order.ItemID,
		Quantity:  order.Quantity,
		Total:     order.Total,
		CreatedAt: order.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *OrderHandler) Get(w http.ResponseWriter, r *http.Request) {
	orderID := r.PathValue("id")
	if orderID == "" {
		ItemRespondWithError(w, http.StatusBadRequest, "Order ID is required to get")
		return
	}

	orderInput := order.GetOrderInput{
		ID: orderID,
	}

	fetchedOrder, err := h.GetOrder.Execute(orderInput) // Changed variable name from 'payment' to 'fetchedItem'
	if err != nil {
		// It's better to check the specific error returned by the use case (e.g., from repository)
		// For now, a generic "item not found" check is used.
		if errors.Is(err, errors.New("Order not found")) { // Changed from specific string comparison to errors.Is
			ItemRespondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		ItemRespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fetchedOrder) // Changed variable name
}
