package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"pabiosoft/dto"
)

func GetArticles(c echo.Context, db *sql.DB) error {
	log.Println("Nouvelle fonction GetArticles appelée")

	// Requête SQL principale pour récupérer les articles
	query := `
        SELECT 
            a.id, a.type, a.cover_image_url, a.cover_text, a.date, 
            a.url, a.profile_image_url, a.estimate_time, 
            a.meta_title, a.meta_description, a.created_at, a.updated_at,a.tags,
            au.id AS author_id, au.name AS author_name, au.country AS author_country, au.profile_image_url AS author_profile_image_url,
            s.id AS status_id, s.name AS status_name,
            v.id AS visibility_id, v.name AS visibility_name
        FROM articles a
        LEFT JOIN authors au ON a.author_id = au.id
        LEFT JOIN statuses s ON a.status_id = s.id
        LEFT JOIN visibilities v ON a.visibility_id = v.id
    `

	// Exécuter la requête pour récupérer les articles
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête SQL : %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erreur lors de la récupération des articles"})
	}
	defer rows.Close()

	var articles []dto.ArticleDTO

	for rows.Next() {
		var article dto.ArticleDTO
		var authorID, authorName, authorCountry, authorProfileImageUrl sql.NullString
		var statusID, statusName, visibilityID, visibilityName sql.NullString
		var tagsJSON sql.NullString

		err := rows.Scan(
			&article.ID, &article.Type, &article.CoverImageUrl, &article.CoverText, &article.Date,
			&article.URL, &article.ProfileImageUrl, &article.EstimateTime,
			&article.MetaTitle, &article.MetaDescription, &article.CreatedAt, &article.UpdatedAt, &tagsJSON,
			&authorID, &authorName, &authorCountry, &authorProfileImageUrl,
			&statusID, &statusName, &visibilityID, &visibilityName,
		)
		if err != nil {
			log.Printf("Erreur lors du Scan des articles : %v", err)
			continue
		}

		// Désérialisez les tags JSON
		if tagsJSON.Valid {
			var tags []string
			if err := json.Unmarshal([]byte(tagsJSON.String), &tags); err == nil {
				article.Tags = tags
			}
		}
		article.Context = "/contexts/Article"
		if authorID.Valid {
			article.Author = &dto.AuthorDTO{
				ID:              authorID.String,
				Name:            authorName.String,
				Country:         authorCountry.String,
				ProfileImageUrl: authorProfileImageUrl.String,
			}
		}
		if statusID.Valid {
			article.Status = &dto.StatusDTO{
				ID:   statusID.String,
				Name: statusName.String,
			}
		}
		if visibilityID.Valid {
			article.Visibility = &dto.VisibilityDTO{
				ID:   visibilityID.String,
				Name: visibilityName.String,
			}
		}

		// Récupérer les chapters et leurs contents
		chapterQuery := `
            SELECT c.id, c.title
            FROM chapters c
            WHERE c.article_id = ?
        `
		chapterRows, err := db.Query(chapterQuery, article.ID)
		if err != nil {
			log.Printf("Erreur lors de la récupération des chapters pour l'article %s : %v", article.ID, err)
		} else {
			var chapters []dto.ChapterDTO
			for chapterRows.Next() {
				var chapter dto.ChapterDTO
				err := chapterRows.Scan(&chapter.ID, &chapter.Title)
				if err != nil {
					log.Printf("Erreur lors du Scan des chapters : %v", err)
					continue
				}

				// Récupérer les contents pour chaque chapter
				contentQuery := `
                    SELECT c.type, c.value, c.language, c.media_type, c.src, c.alt_text
                    FROM contents c
                    WHERE c.chapter_id = ?
                `
				contentRows, err := db.Query(contentQuery, chapter.ID)
				if err != nil {
					log.Printf("Erreur lors de la récupération des contents pour le chapitre %s : %v", chapter.ID, err)
				} else {
					var contents []dto.ContentDTO
					for contentRows.Next() {
						var contentType, contentValue, contentLanguage, contentMediaType, contentSrc, contentAltText sql.NullString
						err := contentRows.Scan(
							&contentType, &contentValue, &contentLanguage,
							&contentMediaType, &contentSrc, &contentAltText,
						)
						if err != nil {
							log.Printf("Erreur lors du Scan des contents : %v", err)
							continue
						}

						content := dto.ContentDTO{
							Type:      contentType.String,
							Value:     contentValue.String,
							Language:  contentLanguage.String,
							MediaType: contentMediaType.String,
							Src:       contentSrc.String,
							AltText:   contentAltText.String,
						}
						contents = append(contents, content)
					}
					contentRows.Close()
					chapter.Content = contents
				}

				chapters = append(chapters, chapter)
			}
			chapterRows.Close()
			article.Chapters = chapters
		}

		// Récupérer les technologies
		technologyQuery := `
            SELECT t.id, t.name, t.logo_url, t.category
            FROM technologies t
            INNER JOIN article_technologies at ON t.id = at.technology_id
            WHERE at.article_id = ?
        `
		technologyRows, err := db.Query(technologyQuery, article.ID)
		if err != nil {
			log.Printf("Erreur lors de la récupération des technologies pour l'article %s : %v", article.ID, err)
		} else {
			var technologies []dto.TechnologyDTO
			for technologyRows.Next() {
				var tech dto.TechnologyDTO
				err := technologyRows.Scan(&tech.ID, &tech.Name, &tech.LogoUrl, &tech.Category)
				if err != nil {
					log.Printf("Erreur lors du Scan des technologies : %v", err)
					continue
				}
				technologies = append(technologies, tech)
			}
			technologyRows.Close()
			article.Technologies = technologies
		}

		// Récupérer les related articles
		relatedArticlesQuery := `
            SELECT ra.related_article_id, a.cover_image_url, au.id AS author_id, au.name, au.country, au.profile_image_url
            FROM related_articles ra
            INNER JOIN articles a ON ra.related_article_id = a.id
            INNER JOIN authors au ON a.author_id = au.id
            WHERE ra.article_id = ?
        `
		relatedRows, err := db.Query(relatedArticlesQuery, article.ID)
		if err != nil {
			log.Printf("Erreur lors de la récupération des related articles pour l'article %s : %v", article.ID, err)
		} else {
			var relatedArticles []dto.RelatedArticleDTO
			for relatedRows.Next() {
				var related dto.RelatedArticleDTO
				var authorID, authorName, authorCountry, authorProfileImageUrl sql.NullString
				err := relatedRows.Scan(
					&related.ID, &related.CoverImageUrl,
					&authorID, &authorName, &authorCountry, &authorProfileImageUrl,
				)
				if err != nil {
					log.Printf("Erreur lors du Scan des related articles : %v", err)
					continue
				}
				related.Author = &dto.AuthorDTO{
					ID:              authorID.String,
					Name:            authorName.String,
					Country:         authorCountry.String,
					ProfileImageUrl: authorProfileImageUrl.String,
				}
				relatedArticles = append(relatedArticles, related)
			}
			relatedRows.Close()
			article.RelatedArticles = relatedArticles
		}

		// Ajouter l'article à la liste finale
		articles = append(articles, article)
	}

	log.Printf("Nombre d'articles récupérés : %d", len(articles))
	return c.JSON(http.StatusOK, articles)
}

//func GetArticlesFactice(c echo.Context) error {
//	var articlesDTO []dto.ArticleDTO
//
//	// Transformer les modèles en DTO
//	for _, article := range models.Articles {
//		articlesDTO = append(articlesDTO, dto.ArticleDTO{
//			Context:       "/contexts/Article",
//			ID:            article.ID,
//			Type:          article.Type,
//			CoverImageUrl: article.CoverImageUrl,
//			CoverText:     article.CoverText,
//			Date:          article.Date,
//			Status: &dto.StatusDTO{ // Transformation pour le champ Status
//				ID:   article.Status.ID,
//				Name: article.Status.Name,
//			},
//			Visibility: &dto.VisibilityDTO{ // Transformation pour le champ Visibility
//				ID:   article.Visibility.ID,
//				Name: article.Visibility.Name,
//			},
//			URL: article.URL,
//			Author: &dto.AuthorDTO{
//				ID:              article.Author.ID,
//				Name:            article.Author.Name,
//				Country:         article.Author.Country,
//				ProfileImageUrl: article.Author.ProfileImageUrl,
//			},
//			Chapters: func() []dto.ChapterDTO {
//				var chapters []dto.ChapterDTO
//				for _, chapter := range article.Chapters {
//					var content []dto.ContentDTO
//					for _, c := range chapter.Content {
//						content = append(content, dto.ContentDTO{
//							Type:      c.Type,
//							Value:     c.Value,
//							Language:  c.Language,
//							MediaType: c.MediaType,
//							Src:       c.Src,
//							AltText:   c.AltText,
//						})
//					}
//					chapters = append(chapters, dto.ChapterDTO{
//						ID:      chapter.ID,
//						Title:   chapter.Title,
//						Content: content,
//					})
//				}
//				return chapters
//			}(),
//			Technologies: func() []dto.TechnologyDTO {
//				var technologies []dto.TechnologyDTO
//				for _, tech := range article.Technologies {
//					technologies = append(technologies, dto.TechnologyDTO{
//						ID:      tech.ID,
//						Name:    tech.Name,
//						LogoUrl: tech.LogoUrl,
//					})
//				}
//				return technologies
//			}(),
//			RelatedArticles: func() []dto.RelatedArticleDTO {
//				var related []dto.RelatedArticleDTO
//				for _, r := range article.RelatedArticles {
//					related = append(related, dto.RelatedArticleDTO{
//						ID:            r.ID,
//						Title:         r.Title,
//						CoverImageUrl: r.CoverImageUrl,
//						Author: &dto.AuthorDTO{
//							ID:      r.Author.ID,
//							Name:    r.Author.Name,
//							Country: r.Author.Country,
//						},
//					})
//				}
//				return related
//			}(),
//			EstimateTime:    article.EstimateTime,
//			MetaTitle:       article.MetaTitle,       // Champ supplémentaire
//			MetaDescription: article.MetaDescription, // Champ supplémentaire
//			CreatedAt:       article.CreatedAt,       // Champ supplémentaire
//			UpdatedAt:       article.UpdatedAt,       // Champ supplémentaire
//		})
//	}
//
//	// Retourner la réponse JSON-LD
//	return c.JSON(http.StatusOK, articlesDTO)
//}
