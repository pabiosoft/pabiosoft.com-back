package handlers

import (
	"net/http"
	"pabiosoft/dto"
	"pabiosoft/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreateArticle(c echo.Context) error {
	// Décoder les données de la requête en DTO
	newArticleDTO := new(dto.ArticleDTO)
	if err := c.Bind(newArticleDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	// Générer un nouvel ID pour l'article
	newID := "/articles/" + uuid.New().String()

	// Convertir le DTO en modèle
	newArticle := models.Article{
		ID:            newID,
		Type:          "Article",
		CoverImageUrl: newArticleDTO.CoverImageUrl,
		CoverText:     newArticleDTO.CoverText,
		Date:          newArticleDTO.Date,
		URL:           newArticleDTO.URL,
		Author: &models.Author{
			ID:              newArticleDTO.Author.ID,
			Name:            newArticleDTO.Author.Name,
			Country:         newArticleDTO.Author.Country,
			ProfileImageUrl: newArticleDTO.Author.ProfileImageUrl,
		},
		Chapters: func() []models.Chapter {
			var chapters []models.Chapter
			for _, chapterDTO := range newArticleDTO.Chapters {
				var content []models.Content
				for _, cDTO := range chapterDTO.Content {
					content = append(content, models.Content{
						Type:      cDTO.Type,
						Value:     cDTO.Value,
						Language:  cDTO.Language,
						MediaType: cDTO.MediaType,
						Src:       cDTO.Src,
						AltText:   cDTO.AltText,
					})
				}
				chapters = append(chapters, models.Chapter{
					ID:      chapterDTO.ID,
					Title:   chapterDTO.Title,
					Content: content,
				})
			}
			return chapters
		}(),
		Technologies: func() []models.Technology {
			var technologies []models.Technology
			for _, techDTO := range newArticleDTO.Technologies {
				technologies = append(technologies, models.Technology{
					ID:      techDTO.ID,
					Name:    techDTO.Name,
					LogoUrl: techDTO.LogoUrl,
				})
			}
			return technologies
		}(),
		RelatedArticles: func() []models.RelatedArticle {
			var relatedArticles []models.RelatedArticle
			for _, relatedDTO := range newArticleDTO.RelatedArticles {
				relatedArticles = append(relatedArticles, models.RelatedArticle{
					ID:            relatedDTO.ID,
					Title:         relatedDTO.Title,
					CoverImageUrl: relatedDTO.CoverImageUrl,
					Author: &models.Author{
						ID:      relatedDTO.Author.ID,
						Name:    relatedDTO.Author.Name,
						Country: relatedDTO.Author.Country,
					},
				})
			}
			return relatedArticles
		}(),
		EstimateTime:    newArticleDTO.EstimateTime,
		MetaTitle:       newArticleDTO.MetaTitle,
		MetaDescription: newArticleDTO.MetaDescription,
		CreatedAt:       newArticleDTO.CreatedAt,
		UpdatedAt:       newArticleDTO.UpdatedAt,
		Status: &models.Status{
			ID:   newArticleDTO.Status.ID,
			Name: newArticleDTO.Status.Name,
		},
		Visibility: &models.Visibility{
			ID:   newArticleDTO.Visibility.ID,
			Name: newArticleDTO.Visibility.Name,
		},
	}

	// Ajouter l'article à la base factice
	models.Articles = append(models.Articles, newArticle)

	// Retourner l'article créé
	return c.JSON(http.StatusCreated, newArticle)
}
