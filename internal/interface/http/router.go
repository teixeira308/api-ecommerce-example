package http

import (
	"net/http"

	"ecommerce-api/internal/interface/http/handler"
)

func NewRouter(
	itemHandler *handler.ItemHandler,
	orderHandler *handler.OrderHandler,
) http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("POST /items", itemHandler.Create)
	mux.HandleFunc("GET /items", itemHandler.List)
	mux.HandleFunc("GET /items/{id}", itemHandler.Get)
	mux.HandleFunc("PUT /items/{id}", itemHandler.Update) // Added update route
	mux.HandleFunc("DELETE /items/{id}", itemHandler.Delete)

	mux.HandleFunc("POST /orders", orderHandler.Create)
	mux.HandleFunc("GET /orders", orderHandler.GetAll)
	mux.HandleFunc("GET /orders/{id}", orderHandler.Get)
	mux.HandleFunc("PUT /orders/{id}", orderHandler.Update)

	return mux
}
