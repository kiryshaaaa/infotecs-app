package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/kiryshaaaa/infotecs-app/internal/service"
)

type Router struct {
	service *service.Service
	Router  *chi.Mux
}

func NewRouter(service *service.Service) *Router {
	r := chi.NewRouter()
	return &Router{service: service, Router: r}
}

func (r *Router) GetHelloWorld() {
	r.Router.Get("/", r.service.MyHandler)
}
