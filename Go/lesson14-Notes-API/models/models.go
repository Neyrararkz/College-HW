package models

type User struct {
	ID       int
	Email    string
	Password string
}

type Note struct {
	ID       int
	Title    string
	Content  string
	IsPublic bool
	UserID   int
}