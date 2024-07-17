package routes

import (
	"backend_article/handlers"
	"backend_article/pkg/mysql"
	"backend_article/repositories"

	"github.com/labstack/echo/v4"
)

func ArticleRoutes(e *echo.Group) {
	articleRepository := repositories.RepositoryArticle(mysql.DB)
	h := handlers.HandlerArticle(articleRepository)

	e.GET("/articles", h.FindArticles)
	e.GET("/article/:id", h.GetArticle)
	e.POST("/article", h.CreateArticle)
}
