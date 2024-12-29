package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

// Fake database
var Users = []User{
	{ID: 1, Name: "dev2", Email: "dev2@mo.com", Password: "123456"},
	{ID: 2, Name: "dev1", Email: "bob@mo.com", Password: "pabios"},
}
