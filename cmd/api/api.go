package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gerardo02/practicando-mas/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, conn *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   conn,
	}
}

func (api *APIServer) Run() error {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	handlers := handlers.NewHandlers(api.db)

	handlers.ManageV1Routes(v1Router)

	router.Mount("/api/v1", v1Router)

	log.Println("test URL: http://localhost:8080/api/v1/auth/google")

	return http.ListenAndServe(api.addr, router)
}
