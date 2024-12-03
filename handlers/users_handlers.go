package handlers

import (
	"net/http"
	"pabiosoft/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Récupérer tous les utilisateurs
func GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, models.Users)
}

// Récupérer un utilisateur par ID
func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, user := range models.Users {
		if user.ID == id {
			return c.JSON(http.StatusOK, user)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
}

func GetUserPost(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, user := range models.Users {
		if user.ID == id {
			// Filtrer les publications écrites par cet utilisateur
			userPosts := []models.Post{}
			for _, post := range models.Posts {
				if post.AuthorID == user.ID {
					userPosts = append(userPosts, post)
				}
			}
			// Ajouter les publications dans la réponse utilisateur
			response := map[string]interface{}{
				"user":  user,
				"posts": userPosts,
			}
			return c.JSON(http.StatusOK, response)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
}

// Créer un nouvel utilisateur
func CreateUser(c echo.Context) error {
	newUser := new(models.User)
	if err := c.Bind(newUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}
	newUser.ID = len(models.Users) + 1
	models.Users = append(models.Users, *newUser)
	return c.JSON(http.StatusCreated, newUser)
}

// Mettre à jour un utilisateur existant
func UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, user := range models.Users {
		if user.ID == id {
			if err := c.Bind(&models.Users[i]); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
			}
			return c.JSON(http.StatusOK, models.Users[i])
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
}

// Supprimer un utilisateur
func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, user := range models.Users {
		if user.ID == id {
			models.Users = append(models.Users[:i], models.Users[i+1:]...)
			return c.JSON(http.StatusOK, map[string]string{"message": "User deleted"})
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
}
