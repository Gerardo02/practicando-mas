package handlers

import (
	"database/sql"

	"github.com/gerardo02/practicando-mas/db"
	"github.com/gerardo02/practicando-mas/services"
)

type Handler struct {
	services *services.Services
	db       *db.Queries
}

func NewHandlers(conn *sql.DB) *Handler {
	return &Handler{
		services: services.NewServices(),
		db:       db.New(conn),
	}
}
