package main

import (
	"remy-aquavelo/config"
	"time"

	"github.com/joho/godotenv"
	"github.com/kataras/golog"
)

// init initializes the application by loading environment variables, initializing the configuration,
// setting the log level, initializing the PostgreSQL connection, and logging the successful initialization.
func init() {
	err := godotenv.Load()
	if err != nil {
		golog.Warn("No .env file found")
	}

	config.Cfg.App = config.InitApp()
	golog.SetLevel(config.Cfg.App.DebugLevel)

	config.Cfg.DB = config.InitPSQL()
	golog.Debug("init successful in " + time.Since(config.Cfg.App.StartedTime).String())
}
