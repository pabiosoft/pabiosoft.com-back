package models

type Post struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID int    `json:"author_id"` // Lien vers l'ID d'un utilisateur

}

// Fake database for Posts
var Posts = []Post{
	{ID: 1, Title: "First Post", Content: "Content of the first post", AuthorID: 1},
	{ID: 2, Title: "Second Post", Content: "Content of the second post", AuthorID: 2},
}
