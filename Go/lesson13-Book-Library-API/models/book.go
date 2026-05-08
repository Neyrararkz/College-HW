package models

type Book struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Year      int    `json:"year"`
	Genre     string `json:"genre"`
	Available bool   `json:"available"`
}