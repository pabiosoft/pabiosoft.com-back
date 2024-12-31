package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pabiosoft/domain/models"
	"pabiosoft/dto"
)

func GetAuthors(c echo.Context) error {
	var authorsDTO []dto.AuthorDTO

	// Transformer les données en DTO
	for _, author := range models.Authors {
		authorsDTO = append(authorsDTO, dto.AuthorDTO{
			ID:              author.ID,
			Name:            author.Name,
			Country:         author.Country,
			ProfileImageUrl: author.ProfileImageUrl,
		})
	}

	// Retourner la liste des auteurs
	return c.JSON(http.StatusOK, authorsDTO)
}
