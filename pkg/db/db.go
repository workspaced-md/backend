package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

type Store struct {
	DB *sql.DB
}

func NewStore() (*Store, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("error %s", err)
	}

	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASS")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("DB connection successful")

	return &Store{
		DB: db,
	}, nil
}

func (s *Store) InitAccountsTable() {
	_, err := s.DB.Exec(`CREATE TABLE IF NOT EXISTS accounts (
		id SERIAL PRIMARY KEY,
		email TEXT NOT NULL,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Fatalf("Error creating accounts table", err)
	}
}
