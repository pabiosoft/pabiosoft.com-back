package handlers

import (
	"net/http"
	"pabiosoft/dto"
	"pabiosoft/models"

	"github.com/labstack/echo/v4"
)

func GetVisibilities(c echo.Context) error {
	var visibilityDTOs []dto.VisibilityDTO
	for _, visibility := range models.Visibilities {
		visibilityDTOs = append(visibilityDTOs, dto.VisibilityDTO{
			ID:   visibility.ID,
			Name: visibility.Name,
		})
	}
	return c.JSON(http.StatusOK, visibilityDTOs)
}
