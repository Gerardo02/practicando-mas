package handlers

import (
	"context"
	"io"
	"log"
	"net/http"
)

func (h *Handler) HandleGoogleOAuthRequest(w http.ResponseWriter, r *http.Request) {
	url := h.services.OAuth.GoogleAuth.AuthCodeURL("randomstate")

	http.Redirect(w, r, url, http.StatusSeeOther)
}

func (h *Handler) HandleOAuthCallback(w http.ResponseWriter, r *http.Request) {
	state := r.URL.Query().Get("state")
	if state != "randomstate" {
		w.Write([]byte("state is not random state"))
		return
	}

	code := r.URL.Query().Get("code")
	h.googleHandler(w, r, code)
	// provider := r.URL.Query().Get("provider")
	// switch provider {
	// case "google":
	// default:
	// 	w.Write([]byte("no provider por aqui compa"))
	// }
}

func (h *Handler) googleHandler(w http.ResponseWriter, r *http.Request, code string) {
	googlecon := h.services.OAuth.GoogleAuth

	token, err := googlecon.Exchange(context.Background(), code)
	if err != nil {
		w.Write([]byte("token loco error"))
		return
	}

	resp, err := http.Get(
		"https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken,
	)
	if err != nil {
		w.Write([]byte("resp on user error"))
		return
	}

	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		w.Write([]byte("error parsing data"))
		return
	}
	log.Println(string(userData))
	http.Redirect(w, r, "http://localhost:5173/dashboard", http.StatusSeeOther)
}
