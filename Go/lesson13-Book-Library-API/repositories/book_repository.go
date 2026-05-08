package repositories

import (
	"lesson13-Book-Library-API/config"
	"lesson13-Book-Library-API/models"
)

func GetAllBooks() ([]models.Book, error) {
	rows, err := config.DB.Query("SELECT id, title, author, year, genre, available FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var b models.Book
		rows.Scan(&b.ID, &b.Title, &b.Author, &b.Year, &b.Genre, &b.Available)
		books = append(books, b)
	}
	return books, nil
}

func GetBookByID(id string) (*models.Book, error) {
	b := &models.Book{}
	query := "SELECT id, title, author, year, genre, available FROM books WHERE id = $1"
	err := config.DB.QueryRow(query, id).Scan(&b.ID, &b.Title, &b.Author, &b.Year, &b.Genre, &b.Available)
	return b, err
}

func CreateBook(b *models.Book) error {
	query := "INSERT INTO books (title, author, year, genre) VALUES ($1, $2, $3, $4) RETURNING id"
	return config.DB.QueryRow(query, b.Title, b.Author, b.Year, b.Genre).Scan(&b.ID)
}

func UpdateBook(id string, b *models.Book) error {
	query := "UPDATE books SET title=$1, author=$2, year=$3, genre=$4, available=$5 WHERE id=$6"
	_, err := config.DB.Exec(query, b.Title, b.Author, b.Year, b.Genre, b.Available, id)
	return err
}

func DeleteBook(id string) error {
	_, err := config.DB.Exec("DELETE FROM books WHERE id = $1", id)
	return err
}