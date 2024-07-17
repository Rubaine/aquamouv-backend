package config

import "gorm.io/gorm"

// Config represents the global configuration for the application.
type Config struct {
	App App
	DB *gorm.DB
}

var Cfg Config