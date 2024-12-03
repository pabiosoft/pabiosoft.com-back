package dto

type PostDTO struct {
	Context string `json:"@context"`
	ID      string `json:"@id"`
	Type    string `json:"@type"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
