package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func NewStorage() (*sql.DB, error) {
	db, err := sql.Open("postgres", getConnectionString())
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getConnectionString() string {
	godotenv.Load()
	psUSER := os.Getenv("POSTGRES_USER")
	psPSW := os.Getenv("POSTGRES_PASSWORD")
	psNAME := os.Getenv("POSTGRES_DB")
	psPORT := os.Getenv("PS_PORT")

	return fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable", psUSER, psPSW, psNAME, psPORT)
}
