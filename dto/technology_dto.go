package dto

type TechnologyDTO struct {
	ID       string `json:"@id"`
	Name     string `json:"name"`
	LogoUrl  string `json:"logoUrl"`
	Category string `json:"category"`
}
