package links

import (
	"database/sql"
	"fmt"
	"testingTask/internal/shorter"
)

type LinksSQLRepo struct {
	DB *sql.DB
}

func NewLinksSQLRepo(db *sql.DB) *LinksSQLRepo {
	return &LinksSQLRepo{
		DB: db,
	}
}

func (repo *LinksSQLRepo) Add(longURL string) (string, error) {
	var shortURL string
	proceed := longURL
	for {
		shortURL = shorter.Shorter(proceed)
		origin, err := repo.Get(shortURL)
		if err != nil {
			_, err := repo.DB.Exec(
				"INSERT INTO links (short_URL, long_URL) VALUES ($su, $lu)",
				shortURL,
				longURL)
			if err != nil {
				return "", err
			}
			break
		} else if origin != longURL {
			proceed = shortURL
		} else {
			break
		}
	}
	return shortURL, nil
}

func (repo *LinksSQLRepo) Get(shortURL string) (string, error) {
	row := repo.DB.QueryRow(
		"SELECT long_URL FROM links WHERE short_URL=$su",
		shortURL)
	var longURL string
	err := row.Scan(&longURL)
	if err != nil {
		return "", fmt.Errorf("db_error")
	}
	return longURL, nil
}
