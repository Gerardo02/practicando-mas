package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gerardo02/practicando-mas/db"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type User struct {
	OAuthID       string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
}

type Claims struct {
	UserOAuthID string `json:"user_oauth_id"`
	jwt.RegisteredClaims
}

func (h *Handler) HandleGoogleOAuthRequest(w http.ResponseWriter, r *http.Request) {
	url := h.services.OAuth.Google.AuthCodeURL("randomstate")

	http.Redirect(w, r, url, http.StatusSeeOther)
}

func (h *Handler) HandleOAuthCallback(w http.ResponseWriter, r *http.Request) {
	state := r.URL.Query().Get("state")
	if state != "randomstate" {
		w.Write([]byte("state is not random state"))
		return
	}

	code := r.URL.Query().Get("code")
	userData, err := h.googleHandler(code)
	if err != nil {
		w.Write([]byte("error authenticating google"))
		return
	}

	tokenString, err := h.sessionJWT(userData, r)
	if err != nil {
		log.Fatal(err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
		Path:     "/",
	})

	redirectURL := os.Getenv("URL_CLIENT") + "dashboard"
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}

func (h *Handler) HandleWhoAmI(w http.ResponseWriter, r *http.Request) {
	type resonseUser struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value

	claims, err := h.validateJWTToken(tokenStr)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	dbUser, err := h.db.GetUser(r.Context(), claims.UserOAuthID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := resonseUser{
		Id:    dbUser.UserOauthID,
		Name:  dbUser.Name,
		Email: dbUser.Email,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) googleHandler(code string) (User, error) {
	var userData User
	googlecon := h.services.OAuth.Google

	token, err := googlecon.Exchange(context.Background(), code)
	if err != nil {
		return userData, err
	}

	resp, err := http.Get(
		"https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken,
	)
	if err != nil {
		return userData, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&userData)
	if err != nil {
		return userData, err
	}

	log.Println(userData.Name)
	return userData, nil
}

func (h *Handler) sessionJWT(userData User, r *http.Request) (string, error) {
	err := h.db.CreateUser(r.Context(), db.CreateUserParams{
		ID:          uuid.New(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		UserOauthID: userData.OAuthID,
		Email:       userData.Email,
		Name:        userData.Name,
	})
	if err != nil {
		return "", errors.New("error line 151")
	}

	claims := &Claims{
		UserOAuthID: userData.OAuthID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokeString, err := jwtToken.SignedString(jwtKey)
	if err != nil {
		return "", errors.New("error line 164")
	}

	return tokeString, nil
}

func (h *Handler) validateJWTToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
