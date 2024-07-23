package config

import (
	"fmt"
	"os"
	"remy-aquavelo/models"

	"github.com/kataras/golog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitPSQL initializes a PostgreSQL connection and creates a users table if it doesn't exist.
// It reads the database connection details from environment variables: DB_USER, DB_PASSWORD, DB_HOST, and DB_NAME.
// It returns a context and a *pgx.ConnConfig.
func InitPSQL() *gorm.DB {

	var host, user, password, dbname, port string

	if env := os.Getenv("DB_USER"); env == "" {
		golog.Fatal("Bad 'DB_USER' environment variable")
	} else {
		user = env
	}

	if env := os.Getenv("DB_PASSWORD"); env == "" {
		golog.Warn("'DB_PASSWORD' environment variable not set")
	} else {
		password = env
	}

	if env := os.Getenv("DB_HOST"); env == "" {
		golog.Fatal("Bad 'DB_HOST' environment variable")
		os.Exit(1)
	} else {
		host = env
	}

	if env := os.Getenv("DB_PORT"); env == "" {
		golog.Fatal("Bad 'DB_PORT' environment variable")
		os.Exit(1)
	} else {
		port = env
	}

	if env := os.Getenv("DB_NAME"); env == "" {
		golog.Fatal("Bad 'DB_NAME' parameter env")
	} else {
		dbname = env
	}

	dsn :=  fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Paris", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		golog.Fatal("failed to connect database")
	}
  
	// Migrate the schema
	db.AutoMigrate(&models.ContactInfo{})

	return db

}
