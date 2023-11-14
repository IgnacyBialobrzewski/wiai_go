package models

import (
	"crypto/sha256"
	"database/sql"
	"errors"
	"fmt"
)

type User struct {
	Id       int64
	Username string
	Password string
}

var ErrUserAlreadyExists error = errors.New("user already exists")
var ErrInvalidCredentials = errors.New("invalid credentials")

const queryCreateUser = `--sql
	INSERT OR IGNORE INTO users (Username, Password)
	VALUES (?, ?)
`
const queryLoadUser = `--sql
	SELECT Id, Username, Password FROM users
	WHERE Username = ? AND Password = ?
`

func CreateUser(db *sql.DB, username string, password string) (*User, error) {
	h := sha256.New()
	_, err := h.Write([]byte(password))

	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	hashedPassword := fmt.Sprintf("%x", h.Sum(nil))
	result, err := db.Exec(queryCreateUser, username, hashedPassword)

	if err != nil {
		return nil, fmt.Errorf("failed to execute stmt: %w", err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return nil, fmt.Errorf("failed to get affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return nil, ErrUserAlreadyExists
	}

	userId, err := result.LastInsertId()

	if err != nil {
		return nil, fmt.Errorf("failed to get inserted id: %w", err)
	}

	return &User{userId, username, hashedPassword}, nil
}

func LoadUser(db *sql.DB, username string, password string) (*User, error) {
	h := sha256.New()
	_, err := h.Write([]byte(password))

	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	hashedPassword := fmt.Sprintf("%x", h.Sum(nil))

	var user User
	err = db.
		QueryRow(queryLoadUser, username, hashedPassword).
		Scan(&user.Id, &user.Username, &user.Password)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrInvalidCredentials
		}

		return nil, fmt.Errorf("failed to copy row to struct: %w", err)
	}

	return &user, nil
}
