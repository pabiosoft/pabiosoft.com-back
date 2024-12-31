package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func UpdateArticleURL(c echo.Context, db *sql.DB) error {
	// Récupérer l'ID de l'article depuis les paramètres
	articleID := c.Param("id")

	// Payload reçu
	var payload struct {
		URL string `json:"url" validate:"required,url"`
	}
	if err := c.Bind(&payload); err != nil {
		log.Printf("Erreur de binding du payload : %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	// Appel à la fonction de mise à jour
	err := updateArticleURLInDB(db, articleID, payload.URL)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Article not found"})
		}
		log.Printf("Erreur lors de la mise à jour de l'URL : %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update article URL"})
	}

	// Réponse de succès
	//return c.JSON(http.StatusOK, map[string]string{"message": "Article URL updated successfully"})
	// Réponse de succès avec le message et l'URL mise à jour
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Article URL updated successfully",
		"url":     payload.URL,
	})
}

func updateArticleURLInDB(db *sql.DB, articleID string, url string) error {
	// Démarrer une transaction
	tx, err := db.Begin()
	if err != nil {
		log.Printf("Erreur lors du démarrage de la transaction : %v", err)
		return fmt.Errorf("erreur interne : %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// Vérification de l'existence de l'article
	var articleExists bool
	err = tx.QueryRow("SELECT EXISTS (SELECT 1 FROM articles WHERE id = ?)", articleID).Scan(&articleExists)
	if err != nil {
		log.Printf("Erreur lors de la vérification de l'article : %v", err)
		return fmt.Errorf("erreur interne : %w", err)
	}
	if !articleExists {
		log.Printf("Article avec l'ID %s non trouvé", articleID)
		return sql.ErrNoRows
	}

	// Mise à jour de l'URL
	query := `UPDATE articles SET url = ? WHERE id = ?`
	result, err := tx.Exec(query, url, articleID)
	if err != nil {
		log.Printf("Erreur lors de la mise à jour de l'URL de l'article : %v", err)
		return fmt.Errorf("erreur interne : %w", err)
	}

	// Vérifie si une ligne a été affectée
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Erreur lors de la vérification des lignes affectées : %v", err)
		return fmt.Errorf("erreur interne : %w", err)
	}
	if rowsAffected == 0 {
		log.Printf("Aucune ligne affectée pour l'ID %s", articleID)
		return sql.ErrNoRows
	}

	log.Printf("URL mise à jour avec succès pour l'article %s : %s", articleID, url)
	return nil
}
