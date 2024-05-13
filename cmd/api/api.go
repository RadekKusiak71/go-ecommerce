package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/RadekKusiak71/goEcom/services/accounts"
	"github.com/RadekKusiak71/goEcom/services/products"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(listenAddr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: listenAddr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()

	accountStore := accounts.NewStore(s.db)
	accountsHandler := accounts.NewHandler(accountStore)
	accountsHandler.RegisterRoutes(router)

	productStore := products.NewStore(s.db)
	productsHandler := products.NewHandler(productStore)
	productsHandler.RegisterRoutes(router)

	log.Printf("Server is running on localhost%s \n", s.addr)
	return http.ListenAndServe(s.addr, router)
}
