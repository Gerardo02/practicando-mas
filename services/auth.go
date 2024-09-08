package services

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type OAuth struct {
	GoogleAuth oauth2.Config
}

// type Config struct {
// 	GoogleLoginConfig oauth2.Config
// }

var AppConfig OAuth

func NewAuth() *OAuth {
	return &OAuth{
		GoogleAuth: GoogleConfig(),
	}
}

func GoogleConfig() oauth2.Config {
	AppConfig.GoogleAuth = oauth2.Config{
		RedirectURL:  "http://localhost:8080/api/v1/auth/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	return AppConfig.GoogleAuth
}
