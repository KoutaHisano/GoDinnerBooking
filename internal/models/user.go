package models

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	LoginID  string `json:"login_id"`
	Password string `json:"password"`
}

func CreateUser(db *sql.DB, u User) error {
	_, err := db.Exec("INSERT INTO users (login_id, password) VALUES (?, ?)", u.LoginID, u.Password)
	return err
}

func GetUser(db *sql.DB, id int) (*User, error) {
	u := &User{}
	err := db.QueryRow("SELECT id, login_id, password FROM users WHERE id = ?", id).Scan(&u.ID, &u.LoginID, &u.Password)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// HashPassword は与えられた平文のパスワードをハッシュ化します。
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
