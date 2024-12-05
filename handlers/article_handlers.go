package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pabiosoft/dto"
	"pabiosoft/models"
)

func GetArticles(c echo.Context) error {
	var articlesDTO []dto.ArticleDTO

	// Transformer les modèles en DTO
	for _, article := range models.Articles {
		articlesDTO = append(articlesDTO, dto.ArticleDTO{
			Context:       "/contexts/Article",
			ID:            article.ID,
			Type:          article.Type,
			CoverImageUrl: article.CoverImageUrl,
			CoverText:     article.CoverText,
			Date:          article.Date,
			Status: &dto.StatusDTO{ // Transformation pour le champ Status
				ID:   article.Status.ID,
				Name: article.Status.Name,
			},
			Visibility: &dto.VisibilityDTO{ // Transformation pour le champ Visibility
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
			MetaTitle:       article.MetaTitle,       // Champ supplémentaire
			MetaDescription: article.MetaDescription, // Champ supplémentaire
			CreatedAt:       article.CreatedAt,       // Champ supplémentaire
			UpdatedAt:       article.UpdatedAt,       // Champ supplémentaire
		})
	}

	// Retourner la réponse JSON-LD
	return c.JSON(http.StatusOK, articlesDTO)
}
