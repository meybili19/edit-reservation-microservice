package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func ConnectDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error verifying connection to database: %w", err)
	}
	return db, nil
}

func InitDatabases() (map[string]*sql.DB, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	databases := map[string]string{
		"reservations": os.Getenv("DB_RESERVATIONS_DSN"),
	}

	for name, dsn := range databases {
		if dsn == "" {
			return nil, fmt.Errorf("missing DSN for %s", name)
		}
	}

	connections := make(map[string]*sql.DB)
	for name, dsn := range databases {
		db, err := ConnectDB(dsn)
		if err != nil {
			return nil, fmt.Errorf("error connecting to %s: %w", name, err)
		}
		connections[name] = db
	}
	return connections, nil
}
