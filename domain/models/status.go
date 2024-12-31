package models

type Status struct {
	ID   string `json:"@id"`
	Name string `json:"name"` // E.g., "draft", "published"
}

var Statuses = []Status{
	{ID: "c1b9b2d1-3f4b-4c5f-9a1c-3e9d6a8f2b78", Name: "Draft"},
	{ID: "d2a9c3d4-6a3f-7c8d-9b1e-4e7f5a9b6e2c", Name: "Published"},
}
