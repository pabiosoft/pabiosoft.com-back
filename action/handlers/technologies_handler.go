package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pabiosoft/domain/models"
	"pabiosoft/dto"
)

func GetTechnologies(c echo.Context) error {
	var technologiesDTO []dto.TechnologyDTO

	// Transformer les données en DTO
	for _, tech := range models.Technologies {
		technologiesDTO = append(technologiesDTO, dto.TechnologyDTO{
			ID:       tech.ID,
			Name:     tech.Name,
			LogoUrl:  tech.LogoUrl,
			Category: tech.Category,
		})
	}

	// Retourner la liste des technologies
	return c.JSON(http.StatusOK, technologiesDTO)
}
