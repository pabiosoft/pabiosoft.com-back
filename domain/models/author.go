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
		ID:              "8e21a1ff-4cd7-4c7c-9394-2c35a7e7a1a2",
		Name:            "BALDE isma",
		Country:         "France",
		ProfileImageUrl: "https://img.icons8.com/?size=100&id=23242&format=png&color=000000",
	},
	{
		ID:              "e12a3aff-5da4-4cd7-8895-2f9b6a9e8c1b",
		Name:            "John Smith",
		Country:         "Belgium",
		ProfileImageUrl: "https://img.icons8.com/?size=100&id=23242&format=png&color=000000",
	},
}
