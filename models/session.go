package models

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Session struct {
	Id     int64
	Key    string
	UserId int64
}

const queryCreateSession = `--sql
	INSERT INTO sessions (Key, UserId)
	VALUES (?, ?)
`
const queryDeleteSession = `--sql
	DELETE FROM sessions
	WHERE Key = ?
`

func CreateSession(db *sql.DB, user *User) (*Session, error) {
	sessionKey := uuid.NewString()
	result, err := db.Exec(queryCreateSession, sessionKey, user.Id)

	if err != nil {
		return nil, fmt.Errorf("failed to insert new session: %w", err)
	}

	sessionId, err := result.LastInsertId()

	if err != nil {
		return nil, fmt.Errorf("failed to get last inserted id: %w", err)
	}

	return &Session{sessionId, sessionKey, user.Id}, nil
}

func (s *Session) Delete(db *sql.DB) error {
	_, err := db.Exec(queryDeleteSession, s.Key)

	if err != nil {
		return fmt.Errorf("failed to delete session row: %w", err)
	}

	return nil
}