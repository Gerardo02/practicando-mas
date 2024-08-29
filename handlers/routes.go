package handlers

import (
	"github.com/go-chi/chi/v5"
)

func (h *Handler) ManageRoutes(router *chi.Mux) {
	router.Get("/orders", h.HandlerGetOrders)
}
