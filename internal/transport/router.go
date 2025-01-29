package transport

import (
	"github.com/go-chi/chi/v5"
	"github.com/kiryshaaaa/infotecs-app/internal/services"
)

type Router struct {
	apihandler *APIHandlers
	Router     *chi.Mux
}

type APIHandlers struct {
	transactionService *services.TransactionService
	walletService      *services.WalletService
}

func NewRouter(transactionService *services.TransactionService, walletService *services.WalletService) *Router {
	apihandler := &APIHandlers{
		transactionService: transactionService,
		walletService:      walletService,
	}
	r := chi.NewRouter()

	router := &Router{apihandler: apihandler, Router: r}

	router.SetupRoutes()

	return router
}

func (r *Router) SetupRoutes() {
	r.Router.Get("/api/transactions", r.apihandler.GetLastHandler)
	r.Router.Get("/api/wallet/{address}/balance", r.apihandler.GetBalanceHandler)
	r.Router.Post("/api/send", r.apihandler.SendHandler)
}
