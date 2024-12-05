package dto

type AuthorDTO struct {
	ID              string `json:"@id"`
	Name            string `json:"name"`
	Country         string `json:"country"`
	ProfileImageUrl string `json:"profileImageUrl"`
}
