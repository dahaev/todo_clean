package repository

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

const (
	userTable      = "users"
	todoListsTable = "todo_lists"
	userListsTable = "users_lists"
	todoItemsTable = "lists_items"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBname   string
	SSLmode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.DBname, cfg.Username, cfg.Password, cfg.SSLmode))
	if err != nil {
		log.Fatal("Error sql open")
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Error ping")
		return nil, err
	}

	return db, nil
}
