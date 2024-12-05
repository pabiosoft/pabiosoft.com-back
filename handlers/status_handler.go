package handlers

import (
	"net/http"
	"pabiosoft/dto"
	"pabiosoft/models"

	"github.com/labstack/echo/v4"
)

func GetStatuses(c echo.Context) error {
	var statusDTOs []dto.StatusDTO
	for _, status := range models.Statuses {
		statusDTOs = append(statusDTOs, dto.StatusDTO{
			ID:   status.ID,
			Name: status.Name,
		})
	}
	return c.JSON(http.StatusOK, statusDTOs)
}
