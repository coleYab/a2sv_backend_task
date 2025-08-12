package utils

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	MongoURI    string
	Email       string
	Port 		string
	AppPassword string
	DatabaseName string
	SecretKey string 
}

var config *Config = nil

func getEnv(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		val = fallback
	}

	return val
}

func NewConfig() Config {
	if config != nil {
		return *config
	}

	newConfig := Config{
		MongoURI:    getEnv("MONGO_URI", ""),
		Email:       getEnv("EMAIL", ""),
		SecretKey: "BiggSecret",
		Port:		 getEnv("API_PORT", ""),
		AppPassword: getEnv("APP_PASSWORD", ""),
		DatabaseName: getEnv("DATABASE_NAME", "blogger"),
	}

	config = &newConfig

	return newConfig
}
