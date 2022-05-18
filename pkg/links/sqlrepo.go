package links

import (
	"database/sql"
	"net/http"
)

type LinksSQLRepo struct {
	DB *sql.DB
}

func NewLinksSQLRepo(db *sql.DB) *LinksSQLRepo {
	return &LinksSQLRepo{
		DB: db,
	}
}

func (repo *LinksSQLRepo) Add(longURL string, r *http.Request) (string, error) {
	// make short link from long one
	var shortURL string // obtained with algorithm
	// maybe check for existence
	_, err := repo.DB.ExecContext(
		r.Context(),
		"INSERT INTO links (`shortURL`, `longURL`) VALUES (&su, $lu)",
		shortURL, longURL)
	if err != nil {
		return "", err
	}
	return shortURL, nil
}

func (repo *LinksSQLRepo) Get(shortURL string, r *http.Request) (string, error) {
	row := repo.DB.QueryRowContext(r.Context(), "SELECT longURL FROM links WHERE shortURL=$sl", shortURL)
	var longURL string
	err := row.Scan(&longURL)
	if err != nil {
		return "", err
	}
	return longURL, nil
}
