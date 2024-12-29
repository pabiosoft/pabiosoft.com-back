package dbchecker

import (
	"database/sql"
	"fmt"
	"github.com/labstack/gommon/log"
)

func EnsureDatabaseExists(db *sql.DB, databaseName string) error {
	// Vérifier si la base existe
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = '%s')", databaseName)
	err := db.QueryRow(query).Scan(&exists)
	if err != nil {
		return fmt.Errorf("erreur lors de la vérification de la base de données : %w", err)
	}

	if exists {
		log.Printf("La base de données '%s' existe déjà.\n", databaseName)
		return nil
	}

	// Créer la base si elle n'existe pas
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", databaseName))
	if err != nil {
		return fmt.Errorf("erreur lors de la création de la base de données : %w", err)
	}

	log.Printf("La base de données '%s' a été créée avec succès.\n", databaseName)
	return nil
}
