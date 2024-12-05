package models

type Visibility struct {
	ID   string `json:"id"`
	Name string `json:"name"` // E.g., "private", "public"
}

var Visibilities = []Visibility{
	{ID: "1", Name: "Private"},
	{ID: "2", Name: "Public"},
}
