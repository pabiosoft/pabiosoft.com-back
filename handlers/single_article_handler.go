package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pabiosoft/dto"
	"pabiosoft/models"
)

func GetSingleArticle(c echo.Context) error {
	id := c.Param("id")

	// Trouver l'article correspondant
	for _, article := range models.Articles {
		if article.ID == "/articles/"+id {
			// Transformer l'article en DTO
			articleDTO := dto.ArticleDTO{
				Context:       "/contexts/Article",
				ID:            article.ID,
				Type:          article.Type,
				CoverImageUrl: article.CoverImageUrl,
				CoverText:     article.CoverText,
				Date:          article.Date,
				Status: &dto.StatusDTO{
					ID:   article.Status.ID,
					Name: article.Status.Name,
				},
				Visibility: &dto.VisibilityDTO{
					ID:   article.Visibility.ID,
					Name: article.Visibility.Name,
				},
				URL: article.URL,
				Author: &dto.AuthorDTO{
					ID:              article.Author.ID,
					Name:            article.Author.Name,
					Country:         article.Author.Country,
					ProfileImageUrl: article.Author.ProfileImageUrl,
				},
				Chapters: func() []dto.ChapterDTO {
					var chapters []dto.ChapterDTO
					for _, chapter := range article.Chapters {
						var content []dto.ContentDTO
						for _, c := range chapter.Content {
							content = append(content, dto.ContentDTO{
								Type:      c.Type,
								Value:     c.Value,
								Language:  c.Language,
								MediaType: c.MediaType,
								Src:       c.Src,
								AltText:   c.AltText,
							})
						}
						chapters = append(chapters, dto.ChapterDTO{
							ID:      chapter.ID,
							Title:   chapter.Title,
							Content: content,
						})
					}
					return chapters
				}(),
				Technologies: func() []dto.TechnologyDTO {
					var technologies []dto.TechnologyDTO
					for _, tech := range article.Technologies {
						technologies = append(technologies, dto.TechnologyDTO{
							ID:      tech.ID,
							Name:    tech.Name,
							LogoUrl: tech.LogoUrl,
						})
					}
					return technologies
				}(),
				RelatedArticles: func() []dto.RelatedArticleDTO {
					var related []dto.RelatedArticleDTO
					for _, r := range article.RelatedArticles {
						related = append(related, dto.RelatedArticleDTO{
							ID:            r.ID,
							Title:         r.Title,
							CoverImageUrl: r.CoverImageUrl,
							Author: &dto.AuthorDTO{
								ID:      r.Author.ID,
								Name:    r.Author.Name,
								Country: r.Author.Country,
							},
						})
					}
					return related
				}(),
				EstimateTime:    article.EstimateTime,
				MetaTitle:       article.MetaTitle,
				MetaDescription: article.MetaDescription,
				CreatedAt:       article.CreatedAt,
				UpdatedAt:       article.UpdatedAt,
			}

			return c.JSON(http.StatusOK, articleDTO)
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{"message": "Article not found"})
}
