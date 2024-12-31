package routes

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
	handlers2 "pabiosoft/action/handlers"
	"pabiosoft/action/utils"
)

func RegisterRoutes(e *echo.Echo, db *sql.DB) {

	e.GET("/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "CORS is working!"})
	})

	// User routes
	e.GET("/users", handlers2.GetUsers)
	e.GET("/users/:id", handlers2.GetUser)
	e.POST("/users", handlers2.CreateUser)
	e.PUT("/users/:id", handlers2.UpdateUser)
	e.DELETE("/users/:id", handlers2.DeleteUser)

	// Post routes
	e.GET("/posts", handlers2.GetPosts)
	e.GET("/posts/:id", handlers2.GetPost)
	e.POST("/posts", handlers2.CreatePost)
	e.PUT("/posts/:id", handlers2.UpdatePost)
	e.DELETE("/posts/:id", handlers2.DeletePost)

	//
	e.GET("/users/:id/posts", handlers2.GetUserWithPostsJSONLD)
	//
	//e.GET("/articles", handlers2.GetArticles)
	e.GET("/articles", func(c echo.Context) error {
		return handlers2.GetArticles(c, db) // Assurez-vous de bien passer `db`
	})
	e.GET("/articles/:id", func(c echo.Context) error {
		return handlers2.GetSingleArticle(c, db) // Assurez-vous de bien passer `db`
	})
	e.POST("/articles", func(c echo.Context) error {
		return handlers2.CreateArticle(c, db) // Assurez-vous de bien passer `db`
	})
	e.PATCH("/articles/:id/url", func(c echo.Context) error {
		return handlers2.UpdateArticleURL(c, db)
	})

	//e.POST("/articles", handlers2.CreateArticle)

	//
	e.GET("/technologies", handlers2.GetTechnologies)
	//
	e.GET("/authors", handlers2.GetAuthors)
	// Status and Visibility routes
	e.GET("/statuses", handlers2.GetStatuses)
	e.GET("/visibilities", handlers2.GetVisibilities)
	//
	e.POST("/upload", handlers2.UploadFileHandler)
	//
	// Route de test
	e.GET("/test-db", func(c echo.Context) error {
		return utils.TestDBConnection(c, db)
	})

	// Servir les fichiers statiques
	e.Static("/uploads", "./public/uploads")

}
