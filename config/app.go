package config

import (
	"os"
	"time"
)

type App struct {
	Name        string
	Version     string
	Port        string
	DebugLevel  string
	StartedTime time.Time
}

// InitApp initializes the application with default values and environment variables.
// It returns an instance of the App struct.
func InitApp() (app App) {
	app.StartedTime = time.Now()

	if env := os.Getenv("APP_NAME"); env != "" {
		app.Name = env
	} else {
		app.Name = "API - Aquavelo"
	}

	if env := os.Getenv("APP_VERSION"); env != "" {
		app.Version = env
	} else {
		app.Version = "v0.0"
	}

	if env := os.Getenv("APP_PORT"); env != "" {
		app.Port = env
	} else {
		app.Port = "4000"
	}

	if env := os.Getenv("APP_LOG_LEVEL"); env != "" {
		app.DebugLevel = env
	} else {
		app.DebugLevel = "debug"
	}

	return
}
