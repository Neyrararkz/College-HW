package models

type Note struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	IsPublic bool `json:"is_public"`
	UserID string `json:"user_id"`
}