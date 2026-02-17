package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"ecommerce-api/internal/infrastructure/config"
	mysqlRepo "ecommerce-api/internal/infrastructure/database/mysql"
	httpRouter "ecommerce-api/internal/interface/http"
	httpHandler "ecommerce-api/internal/interface/http/handler" // Alias for clarity
	itemUsecase "ecommerce-api/internal/usecase/item"           // Alias for clarity
	orderUsecase "ecommerce-api/internal/usecase/order"         // Alias for clarity
)

func main() {

	cfg := config.Load()

	db, err := sql.Open("mysql", cfg.MySQLDSN())
	if err != nil {
		log.Fatal(err)
	}

	itemRepo := mysqlRepo.NewItemRepository(db)
	orderRepo := mysqlRepo.NewOrderRepository(db)

	// Item Use Cases
	createItem := itemUsecase.NewCreateItemUseCase(itemRepo)
	updateItem := itemUsecase.NewUpdateItemUseCase(itemRepo)
	getItem := itemUsecase.NewGetItemUseCase(itemRepo)
	getAllItems := itemUsecase.NewGetAllItemsUseCase(itemRepo)
	deleteItem := itemUsecase.NewDeleteItemUseCase(itemRepo)

	// Order Use Cases
	createOrder := orderUsecase.NewCreateOrderUseCase(orderRepo)
	getAllOrders := orderUsecase.NewGetAllOrdersUseCase(orderRepo)

	// Handlers
	itemHandler := httpHandler.NewItemHandler(
		createItem,
		updateItem,
		getItem,
		getAllItems,
		deleteItem,
	)

	orderHandler := httpHandler.NewOrderHandler(
		createOrder,
		getAllOrders,
	)

	router := httpRouter.NewRouter(
		itemHandler,
		orderHandler,
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
