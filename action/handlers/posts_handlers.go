package handlers

import (
	"net/http"
	models2 "pabiosoft/domain/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Get all posts
func GetPosts(c echo.Context) error {
	return c.JSON(http.StatusOK, models2.Posts)
}

// Get a post by ID
func GetPost(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, post := range models2.Posts {
		if post.ID == id {
			return c.JSON(http.StatusOK, post)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "Post not found"})
}

func CreatePost(c echo.Context) error {
	newPost := new(models2.Post)
	if err := c.Bind(newPost); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	// Vérifier si l'auteur existe
	var authorExists bool
	for _, user := range models2.Users {
		if user.ID == newPost.AuthorID {
			authorExists = true
			break
		}
	}

	if !authorExists {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Author not found"})
	}

	// Ajouter la publication
	newPost.ID = len(models2.Posts) + 1
	models2.Posts = append(models2.Posts, *newPost)

	return c.JSON(http.StatusCreated, newPost)
}

func UpdatePost(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, post := range models2.Posts {
		if post.ID == id {
			if err := c.Bind(&models2.Posts[i]); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
			}
			return c.JSON(http.StatusOK, models2.Posts[i])
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "Post not found"})
}

func DeletePost(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, post := range models2.Posts {
		if post.ID == id {
			// Supprimer la publication de la liste principale
			models2.Posts = append(models2.Posts[:i], models2.Posts[i+1:]...)
			return c.JSON(http.StatusOK, map[string]string{"message": "Post deleted"})
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "Post not found"})
}
