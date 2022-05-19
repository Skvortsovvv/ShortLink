package links

import (
	"crypto/sha256"
)

type Link struct {
	Data string
}

// метод заглушка пока
func Shorter(longURL string) string {
	hash := sha256.Sum256([]byte(longURL))
	result := hash[0:10]
	return string(result)
}

type LinksRepo interface {
	Add(longUrl string) (string, error)
	Get(short string) (string, error)
}
