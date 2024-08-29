package handlers

import (
	"github.com/gerardo02/practicando-mas/db"
	"github.com/gerardo02/practicando-mas/services"
)

type Handler struct {
	services *services.Services
	db       *db.Queries
}

func NewHandlers(services *services.Services, db *db.Queries) *Handler {
	return &Handler{
		services: services,
		db:       db,
	}
}
