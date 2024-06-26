package config

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/kataras/golog"
)

// InitPSQL initializes a PostgreSQL connection and creates a users table if it doesn't exist.
// It reads the database connection details from environment variables: DB_USER, DB_PASSWORD, DB_HOST, and DB_NAME.
// It returns a context and a *pgx.ConnConfig.
func InitPSQL() (context.Context, *pgx.ConnConfig) {
	connstring := "postgres://"

	if env := os.Getenv("DB_USER"); env == "" {
		golog.Fatal("Bad 'DB_USER' environment variable")
	} else {
		connstring += env
	}

	if env := os.Getenv("DB_PASSWORD"); env == "" {
		golog.Warn("'DB_PASSWORD' environment variable not set")
	} else {
		connstring += ":" + env
	}

	if env := os.Getenv("DB_HOST"); env == "" {
		golog.Fatal("Bad 'DB_HOST' environment variable")
		os.Exit(1)
	} else {
		connstring += "@" + env
	}

	if env := os.Getenv("DB_NAME"); env == "" {
		golog.Fatal("Bad 'DB_NAME' parameter env")
	} else {
		connstring += "/" + env
	}

	connstring += "?sslmode=disable"

	connConf, err := pgx.ParseConfig(connstring)
	if err != nil {
		golog.Fatalf("Parse error : %v", err)
	}

	sqlCo, err := pgx.ConnectConfig(context.Background(), connConf)
	if err != nil {
		golog.Errorf("Connection error : %v", err)
		return context.Background(), connConf
	}

	defer sqlCo.Close(context.Background())

	query := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(50) UNIQUE NOT NULL,
			password VARCHAR(50) NOT NULL,
			mail VARCHAR(50) NOT NULL
		);
	`

	_, err = sqlCo.Exec(context.Background(), query)
	if err != nil {
		golog.Errorf("Error creating table : %v", err)
	}
	return context.Background(), connConf

}
