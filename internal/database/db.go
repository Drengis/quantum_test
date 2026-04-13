package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/user/quantum-server/config"
)

var DB *sql.DB

func Init(cfg config.DBConfig) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name,
	)

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	_, err = DB.Exec("CREATE EXTENSION IF NOT EXISTS \"pgcrypto\"")
	if err != nil {
		log.Fatalf("Failed to create extension: %v", err)
	}

	fmt.Printf("Connected to %s database successfully\n", cfg.Name)
	return DB
}
