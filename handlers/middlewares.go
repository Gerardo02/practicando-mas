package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) ManageMiddlewares(router *chi.Mux) {
}

func (h *Handler) AuthMiddleware() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
