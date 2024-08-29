package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gerardo02/practicando-mas/handlers"
	"github.com/go-chi/chi/v5"
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

	v1Router := chi.NewRouter()

	handlers := handlers.NewHandlers(api.db)

	handlers.ManageRoutes(v1Router)

	router.Mount("/api/v1", v1Router)

	log.Println("running server on 8080")

	return http.ListenAndServe(api.addr, router)

}
