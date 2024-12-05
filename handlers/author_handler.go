package handlers

import (
	"net/http"
	"pabiosoft/dto"
	"pabiosoft/models"

	"github.com/labstack/echo/v4"
)

func GetAuthors(c echo.Context) error {
	var authorsDTO []dto.AuthorDTO

	// Transformer les données en DTO
	for _, author := range models.Authors {
		authorsDTO = append(authorsDTO, dto.AuthorDTO{
			ID:              "/authors/" + author.ID,
			Name:            author.Name,
			Country:         author.Country,
			ProfileImageUrl: author.ProfileImageUrl,
		})
	}

	// Retourner la liste des auteurs
	return c.JSON(http.StatusOK, authorsDTO)
}
