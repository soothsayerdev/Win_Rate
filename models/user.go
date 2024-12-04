package models

import (
	"database/sql"
	"errors"
	"winrate/utils"
)

type User struct {
	ID 		 int `json:"id"`
	Email 	 string `json:"email"`
	Password string `json:"password"`
}

func (u *User) Register(db *sql.DB) error {
	// Verify that the email is already registered
	var exists int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", u.Email).Scan(&exists)
	if err != nil || exists > 0 {
		return errors.New("Email already registered")
	}

	// Hash the password for security
	hashedPassword, err := utils.HashPassword(u.Password)
    if err != nil {
        return err
    }

    // Insert the new user into the database
    _, err = db.Exec("INSERT INTO users (email, password) VALUES (?,?)", u.Email, hashedPassword)
    if err != nil {
        return err
    }

    return nil
}

func (u *User) Authenticate(db *sql.DB) (bool, error) {
	var hashedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE email = ?", u.Email).Scan(&hashedPassword)
	if err != nil {
		return false, errors.New("User does not exist")
	}
	return utils.CheckPasswordHash(u.Password, hashedPassword)
}