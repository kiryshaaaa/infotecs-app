package service

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

type Service struct {
	router   *chi.Mux
	database *sql.DB
}

func NewService(router *chi.Mux, database *sql.DB) *Service {
	return &Service{router, database}
}
