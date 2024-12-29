package utils

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

// TestDBConnection teste la connexion à la base de données et vérifie le nombre d'articles dans la table "articles"
func TestDBConnection(c echo.Context, db *sql.DB) error {
	log.Println("TestDBConnection appelée")

	// Vérifier la connexion à la base de données
	err := db.Ping()
	if err != nil {
		log.Println("Échec de la connexion à la base de données")
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Échec de la connexion à la base de données",
			"error":   err.Error(),
		})
	}

	// Vérifier si la table "articles" existe et obtenir le nombre d'articles
	var count int
	countQuery := `SELECT COUNT(*) FROM  articles`
	err = db.QueryRow(countQuery).Scan(&count)
	if err != nil {
		log.Printf("Erreur lors de la requête COUNT : %v\n", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Connexion réussie, mais échec lors de la vérification des articles",
			"error":   err.Error(),
		})
	}

	log.Printf("Nombre d'articles trouvés dans la table articles : %d\n", count)

	// Retourner le résultat
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Connexion à la base de données réussie",
		"count":   count,
	})
}
