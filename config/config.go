package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// LoadEnv carga las variables de entorno desde .env si no están definidas
func LoadEnv() error {
	if os.Getenv("DB_RESERVATIONS_DSN") != "" && os.Getenv("QUERY_RESERVATION_URL") != "" {
		return nil // No cargar .env si ya están definidas
	}

	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}
	return nil
}

// ConnectDB realiza la conexión a MySQL
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

// InitDatabases inicializa las bases de datos y devuelve un mapa de conexiones
func InitDatabases() (map[string]*sql.DB, error) {
	if err := LoadEnv(); err != nil {
		return nil, fmt.Errorf("error loading environment variables: %w", err)
	}

	databases := map[string]string{
		"reservations": os.Getenv("DB_RESERVATIONS_DSN"),
	}

	connections := make(map[string]*sql.DB)
	for name, dsn := range databases {
		if dsn == "" {
			return nil, fmt.Errorf("missing DSN for %s", name)
		}

		db, err := ConnectDB(dsn)
		if err != nil {
			return nil, fmt.Errorf("error connecting to %s: %w", name, err)
		}
		connections[name] = db
	}
	return connections, nil
}
