package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"pabiosoft/dto"
	"strings"

	"github.com/algolia/algoliasearch-client-go/v4/algolia/search"
	"github.com/labstack/echo/v4"
)

// Helper function to build content for indexing
func buildContent(article dto.ArticleDTO) string {
	var contentBuilder strings.Builder

	// Include chapter titles and contents
	for _, chapter := range article.Chapters {
		contentBuilder.WriteString(chapter.Title + " ")
		for _, content := range chapter.Content {
			if content.Type == "text" || content.Type == "code" { // Include both text and code
				contentBuilder.WriteString(content.Value + " ")
			}
		}
	}

	// Include related article titles
	for _, related := range article.RelatedArticles {
		contentBuilder.WriteString(related.Title + " ")
	}

	return contentBuilder.String()
}

// Helper function to extract technology names
func extractTechnologies(technologies []dto.TechnologyDTO) []string {
	var techNames []string
	for _, tech := range technologies {
		techNames = append(techNames, tech.Name)
	}
	return techNames
}

// Function to create a record for Algolia
func createAlgoliaRecord(article dto.ArticleDTO) map[string]interface{} {
	return map[string]interface{}{
		"objectID":        article.ID,
		"title":           article.CoverText,
		"metaDescription": article.MetaDescription,
		"tags":            article.Tags,
		"content":         buildContent(article), // Inclure les contenus texte et code
		"technologies":    extractTechnologies(article.Technologies),
		"author":          article.Author.Name,
		"categories":      article.Type,
		"createdAt":       article.CreatedAt,
		"updatedAt":       article.UpdatedAt,
	}
}
func getAllArticles(db *sql.DB) ([]dto.ArticleDTO, error) {
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

	log.Printf("Nombre d'articles récupérés depuis la base de données : %d", len(articles))
	return articles, nil
}

func SyncAllArticlesToAlgolia(c echo.Context, db *sql.DB, algoliaAppID string, algoliaAPIKey string) error {
	// Initialiser le client Algolia
	client, _ := search.NewClient(algoliaAppID, algoliaAPIKey)

	// Récupérer tous les articles depuis la base de données
	articles, err := getAllArticles(db)
	if err != nil {
		log.Printf("Erreur lors de la récupération des articles : %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erreur lors de la récupération des articles"})
	}
	log.Printf("Nombre d'articles récupérés pour la synchronisation : %d", len(articles))

	// Créer un tableau d'enregistrements pour Algolia
	var records []map[string]interface{}
	for _, article := range articles {
		record := createAlgoliaRecord(article)
		records = append(records, record)
	}

	// Pousser tous les enregistrements vers Algolia
	res, err := client.SaveObjects("articles", records)
	if err != nil {
		log.Printf("Erreur lors de la synchronisation avec Algolia : %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erreur lors de la synchronisation avec Algolia"})
	}

	log.Printf("Tous les articles ont été synchronisés avec Algolia. Task ID: %d", len(res))

	// Réponse de succès
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Tous les articles ont été synchronisés avec Algolia",
	})
}
