package handlers

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	models2 "pabiosoft/domain/models"
	"pabiosoft/dto"
	"time"
)

func CreateArticle(c echo.Context, db *sql.DB) error {
	// Décoder les données de la requête en DTO
	newArticleDTO := new(dto.ArticleDTO)
	if err := c.Bind(newArticleDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	// Générer un nouvel ID pour l'article
	newArticleID := uuid.New().String()

	// Démarrer une transaction
	tx, err := db.Begin()
	if err != nil {
		log.Printf("Erreur lors du démarrage de la transaction : %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// Conversion de la date
	date, err := time.Parse(time.RFC3339, newArticleDTO.Date)
	if err != nil {
		log.Printf("Erreur lors de la conversion de la date : %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid date format"})
	}
	formattedDate := date.Format("2006-01-02 15:04:05")

	// Conversion des champs "createdAt" et "updatedAt"
	createdAt, err := time.Parse(time.RFC3339, newArticleDTO.CreatedAt)
	if err != nil {
		log.Printf("Erreur lors de la conversion de createdAt : %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid createdAt format"})
	}
	formattedCreatedAt := createdAt.Format("2006-01-02 15:04:05")

	updatedAt, err := time.Parse(time.RFC3339, newArticleDTO.UpdatedAt)
	if err != nil {
		log.Printf("Erreur lors de la conversion de updatedAt : %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid updatedAt format"})
	}
	formattedUpdatedAt := updatedAt.Format("2006-01-02 15:04:05")

	log.Printf("ID de l'auteur reçu après binding : %+v", newArticleDTO.Author)
	if newArticleDTO.Author == nil {
		log.Printf("L'objet Author est nil")
	} else {
		log.Printf("ID de l'auteur : %s", newArticleDTO.Author.ID)
	}

	// Vérification de l'existence de l'auteur
	var authorExists bool
	err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM authors WHERE id = ?)", newArticleDTO.Author.ID).Scan(&authorExists)
	if err != nil || !authorExists {
		log.Printf("L'auteur avec l'ID %s n'existe pas", newArticleDTO.Author.ID)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid author ID"})
	}

	if newArticleDTO.Type == "" {
		log.Println("Erreur : Le champ 'type' est vide")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Type cannot be empty"})
	}

	// Insérer l'article dans la table `articles`
	articleQuery := `
        INSERT INTO articles (
            id, type, cover_image_url, cover_text, date, 
            url, profile_image_url, estimate_time, 
            meta_title, meta_description, created_at, updated_at, 
            author_id, status_id, visibility_id
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `
	_, err = tx.Exec(articleQuery,
		newArticleID, newArticleDTO.Type, newArticleDTO.CoverImageUrl, newArticleDTO.CoverText, formattedDate,
		newArticleDTO.URL, newArticleDTO.ProfileImageUrl, newArticleDTO.EstimateTime,
		newArticleDTO.MetaTitle, newArticleDTO.MetaDescription, formattedCreatedAt, formattedUpdatedAt,
		newArticleDTO.Author.ID, newArticleDTO.Status.ID, newArticleDTO.Visibility.ID,
	)
	if err != nil {
		log.Printf("Erreur lors de l'insertion de l'article : %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erreur lors de l'insertion de l'article"})
	}

	// Insérer les chapters et leurs contents
	for _, chapter := range newArticleDTO.Chapters {
		chapterID := uuid.New().String()
		chapterQuery := `
            INSERT INTO chapters (id, article_id, title)
            VALUES (?, ?, ?)
        `
		_, err = tx.Exec(chapterQuery, chapterID, newArticleID, chapter.Title)
		if err != nil {
			log.Printf("Erreur lors de l'insertion du chapitre : %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erreur lors de l'insertion d'un chapitre"})
		}

		// Insérer les contents pour chaque chapter
		for _, content := range chapter.Content {
			contentQuery := `
                INSERT INTO contents (id, chapter_id, type, value, language, media_type, src, alt_text)
                VALUES (?, ?, ?, ?, ?, ?, ?, ?)
            `
			_, err = tx.Exec(contentQuery,
				uuid.New().String(), chapterID, content.Type, content.Value, content.Language,
				content.MediaType, content.Src, content.AltText,
			)
			if err != nil {
				log.Printf("Erreur lors de l'insertion du content : %v", err)
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erreur lors de l'insertion d'un content"})
			}
		}
	}

	// Insérer les technologies
	for _, tech := range newArticleDTO.Technologies {
		technologyQuery := `
            INSERT INTO article_technologies (article_id, technology_id)
            VALUES (?, ?)
        `
		_, err = tx.Exec(technologyQuery, newArticleID, tech.ID)
		if err != nil {
			log.Printf("Erreur lors de l'insertion de la technologie : %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erreur lors de l'insertion d'une technologie"})
		}
	}

	// Vérification et insertion des related articles
	for _, related := range newArticleDTO.RelatedArticles {
		if related.ID == "" {
			log.Printf("L'article lié est manquant ou l'ID est vide")
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid related article ID"})
		}

		// Vérifier l'existence de l'article lié
		var relatedExists bool
		err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM articles WHERE id = ?)", related.ID).Scan(&relatedExists)
		if err != nil || !relatedExists {
			log.Printf("L'article lié avec l'ID %s n'existe pas", related.ID)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid related article ID"})
		}

		// Insérer l'article lié
		relatedQuery := `
        INSERT INTO related_articles (article_id, related_article_id)
        VALUES (?, ?)
    `
		_, err = tx.Exec(relatedQuery, newArticleID, related.ID)
		if err != nil {
			log.Printf("Erreur lors de l'insertion de l'article lié avec ID %s : %v", related.ID, err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erreur lors de l'insertion d'un article lié"})
		}
	}

	// Construire la réponse simplifiée
	response := map[string]interface{}{
		"message": "Article successfully created",
		"data": map[string]interface{}{
			"@context":  "/contexts/Article",
			"@id":       newArticleID,
			"@type":     newArticleDTO.Type,
			"coverText": newArticleDTO.CoverText,
		},
	}

	// Retourner la réponse simplifiée
	return c.JSON(http.StatusCreated, response)
}

func CreateArticleFactice(c echo.Context) error {
	// Décoder les données de la requête en DTO
	newArticleDTO := new(dto.ArticleDTO)
	if err := c.Bind(newArticleDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	// Générer un nouvel ID pour l'article
	newID := "/articles/" + uuid.New().String()

	// Convertir le DTO en modèle
	newArticle := models2.Article{
		ID:            newID,
		Type:          "Article",
		CoverImageUrl: newArticleDTO.CoverImageUrl,
		CoverText:     newArticleDTO.CoverText,
		Date:          newArticleDTO.Date,
		URL:           newArticleDTO.URL,
		Author: &models2.Author{
			ID:              newArticleDTO.Author.ID,
			Name:            newArticleDTO.Author.Name,
			Country:         newArticleDTO.Author.Country,
			ProfileImageUrl: newArticleDTO.Author.ProfileImageUrl,
		},
		Chapters: func() []models2.Chapter {
			var chapters []models2.Chapter
			for _, chapterDTO := range newArticleDTO.Chapters {
				var content []models2.Content
				for _, cDTO := range chapterDTO.Content {
					content = append(content, models2.Content{
						Type:      cDTO.Type,
						Value:     cDTO.Value,
						Language:  cDTO.Language,
						MediaType: cDTO.MediaType,
						Src:       cDTO.Src,
						AltText:   cDTO.AltText,
					})
				}
				chapters = append(chapters, models2.Chapter{
					ID:      chapterDTO.ID,
					Title:   chapterDTO.Title,
					Content: content,
				})
			}
			return chapters
		}(),
		Technologies: func() []models2.Technology {
			var technologies []models2.Technology
			for _, techDTO := range newArticleDTO.Technologies {
				technologies = append(technologies, models2.Technology{
					ID:      techDTO.ID,
					Name:    techDTO.Name,
					LogoUrl: techDTO.LogoUrl,
				})
			}
			return technologies
		}(),
		RelatedArticles: func() []models2.RelatedArticle {
			var relatedArticles []models2.RelatedArticle
			for _, relatedDTO := range newArticleDTO.RelatedArticles {
				relatedArticles = append(relatedArticles, models2.RelatedArticle{
					ID:            relatedDTO.ID,
					Title:         relatedDTO.Title,
					CoverImageUrl: relatedDTO.CoverImageUrl,
					Author: &models2.Author{
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
		Status: &models2.Status{
			ID:   newArticleDTO.Status.ID,
			Name: newArticleDTO.Status.Name,
		},
		Visibility: &models2.Visibility{
			ID:   newArticleDTO.Visibility.ID,
			Name: newArticleDTO.Visibility.Name,
		},
	}

	// Ajouter l'article à la base factice
	models2.Articles = append(models2.Articles, newArticle)

	// Retourner l'article créé
	return c.JSON(http.StatusCreated, newArticle)
}
