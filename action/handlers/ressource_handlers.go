package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	models2 "pabiosoft/domain/models"
	"pabiosoft/dto"
	"strconv"
)

func GetPostsByUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid user ID"})
	}

	// Trouver l'utilisateur correspondant
	var user *models2.User
	for _, u := range models2.Users {
		if u.ID == id {
			user = &u
			break
		}
	}

	if user == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
	}

	userPosts := []map[string]interface{}{}
	for _, post := range models2.Posts {
		if post.AuthorID == id {
			userPosts = append(userPosts, map[string]interface{}{
				"id":      post.ID,
				"title":   post.Title,
				"content": post.Content,
				"author":  user,
			})
		}
	}

	return c.JSON(http.StatusOK, userPosts)
}

func GetUserWithPostsJSONLD(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid user ID"})
	}

	var user *models2.User
	for _, u := range models2.Users {
		if u.ID == id {
			user = &u
			break
		}
	}

	if user == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
	}

	var userPosts []dto.PostDTO
	for _, post := range models2.Posts {
		if post.AuthorID == id {
			userPosts = append(userPosts, dto.PostDTO{
				Context: "/contexts/Post",
				ID:      "/posts/" + strconv.Itoa(post.ID),
				Type:    "Post",
				Title:   post.Title,
				Content: post.Content,
			})
		}
	}

	userDTO := dto.UserDTO{
		Context: "/contexts/User",
		ID:      "/users/" + strconv.Itoa(user.ID),
		Type:    "User",
		Email:   user.Email,
		Name:    user.Name,
		Posts:   userPosts,
	}

	return c.JSON(http.StatusOK, userDTO)
}
