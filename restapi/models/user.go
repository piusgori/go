package models

import (
	"errors"
	"fmt"

	"example.com/restapi/db"
	"example.com/restapi/utils"
)

type User struct {
	ID       int64
	Email    string`binding:"required"`
	Password string`binding:"required"`
}

func (u User) Save() error {
	query := `INSERT INTO users (email, password) VALUES(?, ?)`
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
	query := `SELECT id, password FROM users WHERE email = ?`

	row := db.DB.QueryRow(query, u.Email)
	var retrievedPassword string
	
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		fmt.Println(err)
		return errors.New("invalid credentials")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	fmt.Println(passwordIsValid)
	if !passwordIsValid {
		return errors.New("invalid credentials")
	}

	return nil
}