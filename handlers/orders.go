package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) HandlerGetOrders(w http.ResponseWriter, r *http.Request) {
	payload, err := h.db.GetOrders(r.Context())
	if err != nil {
		w.Write([]byte("error compa que dices pues"))
		return
	}

	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)
}
