package database

import (
	"database/sql"
	"fmt"

	"go-wsl-project/internal/infrastructure/config"

	_ "github.com/lib/pq"
)

func Connect(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)
	return sql.Open("postgres", dsn)
}
