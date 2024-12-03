package routes

import (
	"pabiosoft/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	// User routes
	e.GET("/users", handlers.GetUsers)
	e.GET("/users/:id", handlers.GetUser)
	e.POST("/users", handlers.CreateUser)
	e.PUT("/users/:id", handlers.UpdateUser)
	e.DELETE("/users/:id", handlers.DeleteUser)

	// Post routes
	e.GET("/posts", handlers.GetPosts)
	e.GET("/posts/:id", handlers.GetPost)
	e.POST("/posts", handlers.CreatePost)
	e.PUT("/posts/:id", handlers.UpdatePost)
	e.DELETE("/posts/:id", handlers.DeletePost)

	//
	e.GET("/users/:id/posts", handlers.GetUserWithPostsJSONLD)
}
