package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"pabiosoft/action/utils"
	"pabiosoft/domain/config"
	"pabiosoft/routes"
)

func main() {
	// Charger la configuration
	cfg := utils.LoadEnv()

	// Connexion admin à MariaDB
	adminDB, err := config.NewMariaDBAdminConnection(cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	if err != nil {
		log.Fatalf("Erreur lors de la connexion admin à MariaDB : %v", err)
	}
	defer func(adminDB *sql.DB) {
		err := adminDB.Close()
		if err != nil {
			log.Printf("Erreur lors de la fermeture de la connexion à la base de données : %v", err)
		}
	}(adminDB)

	log.Printf("La base de données '%s' est prête.\n", cfg.Database)
	log.Printf("Connexion DB : user=%s host=%s port=%s dbname=%s", cfg.User, cfg.Host, cfg.Port, cfg.Database)

	// Initialiser Echo
	e := echo.New()

	// Enregistrer les routes
	routes.RegisterRoutes(e, adminDB)

	// Activer le mode debug
	e.Debug = true

	// Lancer le serveur
	e.Logger.Fatal(e.Start(":8083"))
}
