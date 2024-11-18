package db

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/arnavsurve/workspaced/pkg/shared"
	"github.com/joho/godotenv"
)

type Store struct {
	DB *gorm.DB
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
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	fmt.Println("DB connection successful")

	return &Store{
		DB: db,
	}, nil
}

func (s *Store) InitAccountsTable() {
	err := s.DB.AutoMigrate(&shared.Account{})
	if err != nil {
		log.Fatalf("Error creating accounts table: %v", err)
	}
}
