package models

type Author struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Country         string `json:"country"`
	ProfileImageUrl string `json:"profileImageUrl"`
}

// Fake database for Authors
var Authors = []Author{
	{
		ID:              "23f8b9c4-3f1e-4d8d-992a-3b4f9b4e7a12",
		Name:            "Jane Doe",
		Country:         "France",
		ProfileImageUrl: "https://example.com/images/jane_doe.jpg",
	},
	{
		ID:              "5b9d4c6e-3e2f-4c5f-951b-2a3d4f5b9a78",
		Name:            "Alice Martin",
		Country:         "Belgium",
		ProfileImageUrl: "https://example.com/images/alice_martin.jpg",
	},
	{
		ID:              "8b9c1234-5f4b-4a6d-9f1a-2e7b5f6a8b78",
		Name:            "John Smith",
		Country:         "Germany",
		ProfileImageUrl: "https://example.com/images/john_smith.jpg",
	},
	{
		ID:              "9c1e4f2a-5b4d-6a3f-7e1c-3f9b7a6d8e2c",
		Name:            "Chris Doe",
		Country:         "Netherlands",
		ProfileImageUrl: "https://example.com/images/chris_doe.jpg",
	},
}
