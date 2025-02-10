package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func LoadEnv() error {
	if os.Getenv("DB_RESERVATIONS_HOST") != "" &&
		os.Getenv("DB_RESERVATIONS_USER") != "" &&
		os.Getenv("DB_RESERVATIONS_PASSWORD") != "" &&
		os.Getenv("DB_RESERVATIONS_NAME") != "" &&
		os.Getenv("QUERY_RESERVATION_URL") != "" {
		return nil
	}

	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}
	return nil
}

func GetQueryReservationURL() string {
	return os.Getenv("QUERY_RESERVATION_URL")
}

func ConnectDB(host, user, password, dbname string) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, password, host, dbname)
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

	if err := LoadEnv(); err != nil {
		return nil, fmt.Errorf("error loading environment variables: %w", err)
	}

	dbHost := os.Getenv("DB_RESERVATIONS_HOST")
	dbUser := os.Getenv("DB_RESERVATIONS_USER")
	dbPassword := os.Getenv("DB_RESERVATIONS_PASSWORD")
	dbName := os.Getenv("DB_RESERVATIONS_NAME")

	if dbHost == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		return nil, fmt.Errorf("missing required database environment variables")
	}

	db, err := ConnectDB(dbHost, dbUser, dbPassword, dbName)
	if err != nil {
		return nil, fmt.Errorf("error connecting to reservations database: %w", err)
	}

	return map[string]*sql.DB{"reservations": db}, nil
}
