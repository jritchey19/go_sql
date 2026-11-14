package main

import (
	"fmt"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"modernc.org/sqlite"
	sqlite3 "modernc.org/sqlite/lib"
)

type Store struct {
	store *sql.DB
}

func CreateDB(path string) (Store, error) {
	s := Store{}
	db, err := sql.Open("sqlite", path)
	if err != nil {
		dbErr := fmt.Errorf("createdb: failed to open db: %w", err)
		return Store{}, dbErr
	}

	if err := db.Ping(); err != nil {
		db.Close()
		pingErr := fmt.Errorf("createdb: failed to ping db: %w", err)
		return Store{}, pingErr
	}

	createStmt := `
	CREATE TABLE IF NOT EXISTS users (
		id        TEXT PRIMARY KEY NOT NULL,
		name      TEXT NOT NULL CHECK (name <> '' and LENGTH(name) > 0),
		email     TEXT NOT NULL UNIQUE CHECK (email <> '' and LENGTH(email) > 0)
	);
	`

	_, err = db.Exec(createStmt)
	if err != nil {
		db.Close()
		createTbErr := fmt.Errorf("createdb: failed to create user table: %w", err)
		return Store{}, createTbErr
	}

	s.store = db

	return s, nil
}

func (s *Store) DBClose() {
	s.store.Close()
}

func (s *Store) CreateUser(name, email string) (string, error) {
	newUUID := uuid.New()
	insertStmt := `
	INSERT INTO users (id, name, email) VALUES (?, ?, ?);
	`
	_, err := s.store.Exec(insertStmt, newUUID, name, email)
	if err != nil {
		insertErr := fmt.Errorf("createuser: failed to add user %v, %w", name, err)
		return "", insertErr
	}

	return newUUID.String(), nil
}

func (s *Store) GetUser(id string) (User, error) {
	var user User
	selectStmt := `
	SELECT id, name, email FROM users WHERE id = ?;
	`
	err := s.store.QueryRow(selectStmt, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		selectErr := fmt.Errorf("getuser: failed to get user %v: %w", id, err)
		return User{}, selectErr
	}

	return user, nil
}

func IsEmailUniqueErr (err error) bool {
	var sqliteErr *sqlite.Error
	if errors.As(err, &sqliteErr) {
		return sqliteErr.Code() == sqlite3.SQLITE_CONSTRAINT_UNIQUE
	}
	
	return false
}
