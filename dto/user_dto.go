package dto

type UserDTO struct {
	Context string    `json:"@context"`
	ID      string    `json:"@id"`
	Type    string    `json:"@type"`
	Email   string    `json:"email"`
	Name    string    `json:"name"`
	Posts   []PostDTO `json:"posts"`
}
