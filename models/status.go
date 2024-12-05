package models

type Status struct {
	ID   string `json:"id"`
	Name string `json:"name"` // E.g., "draft", "published"
}

var Statuses = []Status{
	{ID: "1", Name: "Draft"},
	{ID: "2", Name: "Published"},
}
