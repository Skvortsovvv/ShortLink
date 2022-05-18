package links

import "net/http"

type Link struct {
	Data string
}

type LinksRepo interface {
	Add(longUrl string, r *http.Request) (string, error)
	Get(short string, r *http.Request) (string, error)
}
