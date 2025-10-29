package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

// Config holds all configuration for the application.
// Tags are used to map environment variables to struct fields.
type Config struct {
	// Environment string `env:"ENVIRONMENT" envDefault:"development"`
	Server ServerConfig
	// Database    DatabaseConfig
	// JWT         JWTConfig
}

// ServerConfig holds server-specific settings
type ServerConfig struct {
	// Use ":8080" for the port to make it a valid network address
	Port string `env:"SERVER_PORT" envDefault:"8080"`
}

// DatabaseConfig holds database connection settings
// type DatabaseConfig struct {
// 	Host     string `env:"DB_HOST,required"`
// 	Port     int    `env:"DB_PORT,required"`
// 	User     string `env:"DB_USER,required"`
// 	Password string `env:"DB_PASS,required"`
// 	Name     string `env:"DB_NAME,required"`
// }

// JWTConfig holds settings for JWT (JSON Web Tokens)
// type JWTConfig struct {
// 	Secret string `env:"JWT_SECRET,required"`
// }

// Load reads configuration from environment variables and returns a Config struct.
func Load() (*Config, error) {
	// ---
	// Load .env file for local development
	// In production, environment variables are set directly.
	// We ignore the "file not found" error because it's optional.
	// ---
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, loading from environment variables.")
	}

	cfg := &Config{}

	// ---
	// Parse environment variables into the cfg struct
	// The `env:"...` tags tell the library which env var to use.
	// ---
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return cfg, nil
}
