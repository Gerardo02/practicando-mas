package handlers

import (
	"github.com/go-chi/chi/v5"
)

func (h *Handler) ManageV1Routes(router *chi.Mux) {
	router.Get("/orders", h.HandlerGetOrders)
	router.Get("/auth/google", h.HandleGoogleOAuthRequest)
	router.Get("/auth/callback", h.HandleOAuthCallback)
}
