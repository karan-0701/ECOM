package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBName                 string
	JWTExpirationInSeconds int
	JWTSecret              string
}

// Global variable used so that the init config is not called everytime
var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "mypassword"),
		// sprintf is used to format the string
		DBAddress:              fmt.Sprintf("%s:%s", getEnv("DB_HOSTS", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:                 getEnv("DB_NAME", "ecom"),
		JWTExpirationInSeconds: getEnvInt("JWT_EXP", 3600*24*7),
		JWTSecret:              getEnv("JWT_SECRET", "not-secret-secret-anymore?"),
	}
}

// checks if the environment variable exists and if it
// does returns its values
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		// convert the key into int to return the value
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}

	return fallback
}
