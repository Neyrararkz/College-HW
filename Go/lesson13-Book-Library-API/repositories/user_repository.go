package repositories

import (
	"lesson13-Book-Library-API/config"
	"lesson13-Book-Library-API/models"
)

func CreateUser(user *models.User) error {
	query := "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id"
	return config.DB.QueryRow(query, user.Email, user.Password).Scan(&user.ID)
}

func GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := "SELECT id, email, password FROM users WHERE email = $1"
	err := config.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)
	return user, err
}

