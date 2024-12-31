package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"net/http"
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

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:4200"}, // Autorise uniquement cette origine
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions},
		AllowCredentials: true, // Si tu envoies des cookies ou des sessions côté client
		MaxAge:           3600, // Cache les résultats préflight pour 1 heure
	}))

	// Enregistrer les routes
	routes.RegisterRoutes(e, adminDB)

	// Activer le mode debug
	e.Debug = true

	e.Use(middleware.Logger())

	// Lancer le serveur
	e.Logger.Fatal(e.Start(":8083"))
}
