package models

type Content struct {
	Type      string `json:"type"` // text, code, media
	Value     string `json:"value,omitempty"`
	Language  string `json:"language,omitempty"`
	MediaType string `json:"mediaType,omitempty"` // image, video
	Src       string `json:"src,omitempty"`       // URL du média
	AltText   string `json:"altText,omitempty"`
}

type Chapter struct {
	ID      string    `json:"@id"`
	Title   string    `json:"title"`
	Content []Content `json:"content"`
}

type RelatedArticle struct {
	ID            string  `json:"@id"`
	Title         string  `json:"title"`
	CoverImageUrl string  `json:"coverImageUrl"`
	Author        *Author `json:"author"`
}

type Article struct {
	ID              string           `json:"@id"`
	Type            string           `json:"@type"`
	CoverImageUrl   string           `json:"coverImageUrl"`
	CoverText       string           `json:"coverText"`
	Date            string           `json:"date"`
	URL             string           `json:"url"`
	Author          *Author          `json:"author"`
	Chapters        []Chapter        `json:"chapters"`
	Technologies    []Technology     `json:"technologies"`
	ProfileImageUrl string           `json:"profileImageUrl"`
	RelatedArticles []RelatedArticle `json:"relatedArticles"`
	EstimateTime    int              `json:"estimateTime"`
	MetaTitle       string           `json:"metaTitle,omitempty"`
	MetaDescription string           `json:"metaDescription,omitempty"`
	CreatedAt       string           `json:"createdAt,omitempty"`
	UpdatedAt       string           `json:"updatedAt,omitempty"`
	Status          *Status          `json:"status"`
	Visibility      *Visibility      `json:"visibility"`
}

// FACTICES

var Articles = []Article{
	{
		ID:            "/articles/1",
		Type:          "Article",
		CoverImageUrl: "https://img.icons8.com/color/48/vue-js.png",
		CoverText:     "Exploring Vue 3 Composition API for Scalable Applications",
		Date:          "2024-10-25T14:15:00.000Z",
		Status: &Status{
			ID:   "/statuses/published",
			Name: "Published",
		},
		Visibility: &Visibility{
			ID:   "/visibilities/public",
			Name: "Public",
		},
		URL: "https://example.com/vue-composition-api",
		Author: &Author{
			ID:              "/authors/1",
			Name:            "Emily Clarke",
			Country:         "Canada",
			ProfileImageUrl: "https://example.com/images/emily.jpg",
		},
		Chapters: []Chapter{
			{
				ID:    "/chapters/1",
				Title: "Introduction to Composition API",
				Content: []Content{
					{Type: "text", Value: "Vue 3's Composition API offers a flexible way to organize and scale your code."},
					{Type: "code", Language: "javascript", Value: "export default { setup() { return {}; } };"},
					{Type: "media", MediaType: "image", Src: "https://example.com/images/intro.png", AltText: "Intro image"},
				},
			},
		},
		Technologies: []Technology{
			{ID: "/technologies/1", Name: "Vue", LogoUrl: "https://img.icons8.com/color/48/vue-js.png"},
			{ID: "/technologies/2", Name: "JavaScript", LogoUrl: "https://img.icons8.com/color/48/javascript.png"},
		},
		RelatedArticles: []RelatedArticle{
			{
				ID:            "/articles/2",
				Title:         "Using Vue 3 with TypeScript",
				CoverImageUrl: "https://img.icons8.com/color/48/typescript.png",
				Author: &Author{
					ID:      "/authors/2",
					Name:    "Liam Smith",
					Country: "United States",
				},
			},
		},
		EstimateTime:    30,
		MetaTitle:       "Exploring Vue 3 Composition API",
		MetaDescription: "An in-depth exploration of the Vue 3 Composition API, its features, and benefits for scalable applications.",
		CreatedAt:       "2024-10-01T10:00:00.000Z",
		UpdatedAt:       "2024-10-15T12:00:00.000Z",
	},
	{
		ID:            "/articles/2",
		Type:          "Article",
		CoverImageUrl: "https://img.icons8.com/color/48/javascript.png",
		CoverText:     "Getting Started with JavaScript ES6 Features",
		Date:          "2024-11-10T14:15:00.000Z",
		Status: &Status{
			ID:   "/statuses/published",
			Name: "Published",
		},
		Visibility: &Visibility{
			ID:   "/visibilities/private",
			Name: "Private",
		},
		URL: "https://example.com/js-es6-features",
		Author: &Author{
			ID:              "/authors/3",
			Name:            "John Doe",
			Country:         "United Kingdom",
			ProfileImageUrl: "https://example.com/images/john.jpg",
		},
		Chapters: []Chapter{
			{
				ID:    "/chapters/2",
				Title: "Understanding Arrow Functions",
				Content: []Content{
					{Type: "text", Value: "Arrow functions provide a concise way to write function expressions."},
					{Type: "code", Language: "javascript", Value: "const add = (a, b) => a + b;"},
				},
			},
		},
		Technologies: []Technology{
			{ID: "/technologies/2", Name: "JavaScript", LogoUrl: "https://img.icons8.com/color/48/javascript.png"},
		},
		RelatedArticles: []RelatedArticle{
			{
				ID:            "/articles/1",
				Title:         "Exploring Vue 3 Composition API for Scalable Applications",
				CoverImageUrl: "https://img.icons8.com/color/48/vue-js.png",
				Author: &Author{
					ID:      "/authors/1",
					Name:    "Emily Clarke",
					Country: "Canada",
				},
			},
		},
		EstimateTime:    20,
		MetaTitle:       "Getting Started with JavaScript ES6",
		MetaDescription: "Learn about key ES6 features in JavaScript, including arrow functions, template literals, and more.",
		CreatedAt:       "2024-11-01T08:00:00.000Z",
		UpdatedAt:       "2024-11-12T09:30:00.000Z",
	},
}
