package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pabiosoft/domain/models"
	"pabiosoft/dto"
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
