package models

import (
	"errors"

	"example.com/rest_api/db"
	"example.com/rest_api/utils"
)

type User struct {
	ID int64 `json:"id"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)
	var retrievePassword string
	err := row.Scan((&retrievePassword))
	if err != nil {
		return errors.New("Credentials invalid")
	}
	passwordIsValid := utils.CheckPasswordWithHash(u.Password, retrievePassword)
	if !passwordIsValid {
		return errors.New("Credentials invalid")
	}

	return nil
}
