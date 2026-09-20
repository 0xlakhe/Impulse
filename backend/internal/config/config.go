package config

import "os"

type Config struct {
	Port        string
	DatabaseURL string
	JWTSecret   string
	AI          AIConfig
}

type AIConfig struct {
	APIKey  string
	Model   string
	BaseURL string
}

func Load() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	databaseURL := os.Getenv("DATABASE_URL")
	jwtSecret := os.Getenv("JWT_SECRET")
	apiKey := os.Getenv("API_KEY")
	model := os.Getenv("MODEL")
	baseURL := os.Getenv("BASE_URL")

	aiconfig := AIConfig{
		APIKey:  apiKey,
		Model:   model,
		BaseURL: baseURL,
	}
	return &Config{
		Port:        port,
		DatabaseURL: databaseURL,
		JWTSecret:   jwtSecret,
		AI:          aiconfig,
	}
}
