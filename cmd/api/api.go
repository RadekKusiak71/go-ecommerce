package api

import (
	"database/sql"
	"log"
	"net/http"

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
	log.Printf("Server is running on localhost%s \n", s.addr)
	return http.ListenAndServe(s.addr, router)
}
