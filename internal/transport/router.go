package transport

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type Router struct {
	apihandler *APIHandlers
	Router     *chi.Mux
}

type APIHandlers struct{}

func NewRouter() *Router {
	apihandler := APIHandlers{}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	router := &Router{apihandler: &apihandler, Router: r}
	router.GetHelloWorld()
	router.GetLast()
	return router
}

func (r *Router) GetHelloWorld() {
	r.Router.Get("/", r.apihandler.MyHandler)
}

func (r *Router) GetLast() {
	r.Router.Get("/api/transactions", r.apihandler.GetLastHandler)
}

func (r *Router) GetBalance() {

}

func (r *Router) Send() {

}
