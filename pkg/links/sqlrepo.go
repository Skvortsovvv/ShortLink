package links

import (
	"database/sql"
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
				"INSERT INTO links (short_URL, long_URL) VALUES ($1, $2)",
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
		"SELECT long_URL FROM links WHERE short_URL=$1",
		shortURL)
	var longURL string
	err := row.Scan(&longURL)
	if err != nil {
		return "", err
	}
	return longURL, nil
}
