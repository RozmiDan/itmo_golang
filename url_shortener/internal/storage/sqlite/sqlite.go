package sqlite

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/RozmiDan/url_shortener/internal/storage"
	"github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(dbPath string) (*Storage, error) {
	const op = "storage.sqlite.New"

	newDb, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := newDb.Prepare(`
		CREATE TABLE IF NOT EXISTS url(
			id INTEGER PRIMARY KEY,
			alias TEXT NOT NULL UNIQUE,
			url TEXT NOT NULL);
		)
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: newDb}, nil
}

func (s *Storage) SaveURL(urlToSave string, alias string) (int64, error) {
	const op = "storage.sqlite.SaveURL"

	query := `
		INSERT INTO url(alias, url)
		VALUES($1, $2)
		RETURNING id
	`

	res, err := s.db.Exec(query, alias, urlToSave)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok {
			if sqliteErr.Code == sqlite3.ErrConstraint {
				return 0, storage.ErrURLExists
			}
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	resId, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return resId, nil
}

func (s *Storage) GetURL(alias string) (string, error) {
	const op = "storage.sqlite.GetURL"

	query := `
		SELECT url.url FROM url 
		WHERE alias = $1
	`
	var result string

	err := s.db.QueryRow(query, alias).Scan(&result)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", storage.ErrURLNotFound
		}
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return result, nil
}

func (s *Storage) DeleteURL(alias string) error {
	//TODO
	return nil
}
