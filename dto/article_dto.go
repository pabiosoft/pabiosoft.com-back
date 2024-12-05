package dto

type ArticleDTO struct {
	Context         string              `json:"@context"`
	ID              string              `json:"@id"`
	Type            string              `json:"@type"`
	CoverImageUrl   string              `json:"coverImageUrl"`
	CoverText       string              `json:"coverText"`
	Date            string              `json:"date"`
	URL             string              `json:"url"`
	Author          *AuthorDTO          `json:"author"`
	Chapters        []ChapterDTO        `json:"chapters"`
	Technologies    []TechnologyDTO     `json:"technologies"`
	RelatedArticles []RelatedArticleDTO `json:"relatedArticles"`
	EstimateTime    int                 `json:"estimateTime"`
	MetaTitle       string              `json:"metaTitle,omitempty"`
	MetaDescription string              `json:"metaDescription,omitempty"`
	CreatedAt       string              `json:"createdAt,omitempty"`
	UpdatedAt       string              `json:"updatedAt,omitempty"`
	Status          *StatusDTO          `json:"status"`
	Visibility      *VisibilityDTO      `json:"visibility"`
}

type ChapterDTO struct {
	ID      string       `json:"@id"`
	Title   string       `json:"title"`
	Content []ContentDTO `json:"content"`
}

type ContentDTO struct {
	Type      string `json:"type"`
	Value     string `json:"value,omitempty"`
	Language  string `json:"language,omitempty"`
	MediaType string `json:"mediaType,omitempty"`
	Src       string `json:"src,omitempty"`
	AltText   string `json:"altText,omitempty"`
}

type RelatedArticleDTO struct {
	ID            string     `json:"@id"`
	Title         string     `json:"title"`
	CoverImageUrl string     `json:"coverImageUrl"`
	Author        *AuthorDTO `json:"author"`
}
