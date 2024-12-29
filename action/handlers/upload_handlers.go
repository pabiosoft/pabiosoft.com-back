package handlers

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func UploadFileHandler(c echo.Context) error {
	// Obtenez le fichier à partir de la requête
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "File is required"})
	}

	// Ouvrir le fichier pour la lecture
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Unable to open file"})
	}
	defer src.Close()

	// Définir le chemin de sauvegarde
	uploadsDir := "./public/uploads"
	if _, err := os.Stat(uploadsDir); os.IsNotExist(err) {
		err := os.Mkdir(uploadsDir, os.ModePerm)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Unable to create uploads directory"})
		}
	}

	// Générer un nom de fichier unique (UUID) avec l'extension d'origine
	uniqueFilename := fmt.Sprintf("%s%s", uuid.NewString(), filepath.Ext(file.Filename))
	filePath := filepath.Join(uploadsDir, uniqueFilename)

	// Créer le fichier sur le serveur
	dst, err := os.Create(filePath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Unable to save file"})
	}
	defer dst.Close()

	// Copier le contenu du fichier uploadé vers le fichier de destination
	if _, err := io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Unable to copy file content"})
	}

	// Retourner l'URL du fichier sauvegardé
	return c.JSON(http.StatusOK, map[string]string{
		"message": "File uploaded successfully",
		"url":     fmt.Sprintf("/%s/%s", "uploads", uniqueFilename), // URL relative pour servir le fichier
	})
}
