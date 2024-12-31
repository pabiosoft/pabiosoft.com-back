package models

type Visibility struct {
	ID   string `json:"@id"`
	Name string `json:"name"` // E.g., "private", "public"
}

var Visibilities = []Visibility{
	{ID: "a1c3b2d4-3f6a-7c5f-9a1c-4e7f5a9b6e2c", Name: "Private"},
	{ID: "b2d4c3a6-5f7a-9c1e-3f4b-7a8e6c9b2d1f", Name: "Public"},
}
