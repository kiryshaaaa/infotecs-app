package main

import (
	"log"
	"net/http"

	"github.com/kiryshaaaa/infotecs-app/internal/repository"
	"github.com/kiryshaaaa/infotecs-app/internal/services"
	"github.com/kiryshaaaa/infotecs-app/internal/transport"
)

func main() {
	dbStorage, err := repository.NewStorage()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := dbStorage.Close(); err != nil {
			log.Fatalf("Failed to close database connection: %v", err)
		}
	}()

	transactionService := services.NewTransactionService(dbStorage)
	walletService := services.NewWalletService(dbStorage)

	router := transport.NewRouter(transactionService, walletService)
	http.ListenAndServe(":8080", router.Router)
}
