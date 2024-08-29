package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gerardo02/practicando-mas/db"
	"github.com/gerardo02/practicando-mas/handlers"
	"github.com/gerardo02/practicando-mas/services"
	"github.com/go-chi/chi/v5"

	_ "github.com/lib/pq"
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
	database := db.New(api.db)
	router := chi.NewRouter()

	v1Router := chi.NewRouter()

	services := services.NewServices()
	handlers := handlers.NewHandlers(services, database)

	handlers.ManageRoutes(v1Router)

	router.Mount("/api/v1", v1Router)

	log.Println("running server on 8080")

	return http.ListenAndServe(api.addr, router)

}
