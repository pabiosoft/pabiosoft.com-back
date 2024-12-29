package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

// LoadEnv charge les variables d'environnement depuis un fichier .env ou les variables système
func LoadEnv() DBConfig {
	err := godotenv.Load()
	if err != nil {
		log.Println("Pas de fichier .env trouvé, utilisation des variables d'environnement système.")
	}

	return DBConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_NAME"),
	}
}
