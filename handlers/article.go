package handlers

import (
	dto_article "backend_article/dto/article"
	dto "backend_article/dto/result"
	"backend_article/models"
	"backend_article/repositories"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerArticle struct {
	ArticleRepository repositories.ArticleRepository
}

func HandlerArticle(ArticleRepository repositories.ArticleRepository) *handlerArticle {
	return &handlerArticle{ArticleRepository}
}

func (h *handlerArticle) FindArticles(c echo.Context) error {
	search := c.QueryParam("search")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	// page, err := strconv.Atoi(c.QueryParam("page"))
	// if err != nil || page <= 0 {
	// 	page = 1
	// }

	// limit, err := strconv.Atoi(c.QueryParam("limit"))
	// if err != nil || limit <= 0 {
	// 	limit = 10 // Default limit
	// }

	log.Printf("Search: %s, Page: %d, Limit: %d", search, page, limit)

	articles, totalPages, err := h.ArticleRepository.FindArticle(page, limit, search)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	response := map[string]interface{}{
		"articles":    articles,
		"totalPages":  totalPages,
		"currentPage": page,
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: response})
}

func (h *handlerArticle) GetArticle(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var article models.Article
	article, err := h.ArticleRepository.GetArticle(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: article})
}

func (h *handlerArticle) CreateArticle(c echo.Context) error {
	request := new(dto_article.CreateArticleRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	article := models.Article{
		Title:       request.Title,
		Description: request.Description,
		Image:       request.Image,
		PostedAt:    request.PostedAt,
		Creator:     request.Creator,
	}

	article, err = h.ArticleRepository.CreateArticle(article)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: article})
}
