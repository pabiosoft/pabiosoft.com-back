package handlers

import (
	"net/http"
	"pabiosoft/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Get all posts
func GetPosts(c echo.Context) error {
	return c.JSON(http.StatusOK, models.Posts)
}

// Get a post by ID
func GetPost(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, post := range models.Posts {
		if post.ID == id {
			return c.JSON(http.StatusOK, post)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "Post not found"})
}

func CreatePost(c echo.Context) error {
	newPost := new(models.Post)
	if err := c.Bind(newPost); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	// Vérifier si l'auteur existe
	var authorExists bool
	for _, user := range models.Users {
		if user.ID == newPost.AuthorID {
			authorExists = true
			break
		}
	}

	if !authorExists {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Author not found"})
	}

	// Ajouter la publication
	newPost.ID = len(models.Posts) + 1
	models.Posts = append(models.Posts, *newPost)

	return c.JSON(http.StatusCreated, newPost)
}

func UpdatePost(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, post := range models.Posts {
		if post.ID == id {
			if err := c.Bind(&models.Posts[i]); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
			}
			return c.JSON(http.StatusOK, models.Posts[i])
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "Post not found"})
}

func DeletePost(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, post := range models.Posts {
		if post.ID == id {
			// Supprimer la publication de la liste principale
			models.Posts = append(models.Posts[:i], models.Posts[i+1:]...)
			return c.JSON(http.StatusOK, map[string]string{"message": "Post deleted"})
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "Post not found"})
}
