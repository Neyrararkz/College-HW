package repository

import "notes-api/config"

func CreateUser(email string, hashedPassword string) error{
	query := "INSERT INTO users(email,password) VALUES ($1,$2)"
	_,err := config.DB.Exec(query, email,hashedPassword)
	return err
}
func GetUserByEmail(email string)(int,string,error){
	var id int
	var password string

	query := "SELECT id,password FROM users WHERE email = $1"
	err := config.DB.QueryRow(query, email).Scan(&id, &password)

	return id,password,err
}