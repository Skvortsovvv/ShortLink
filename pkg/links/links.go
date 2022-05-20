package links

type Link struct {
	Data string
}

type LinksRepo interface {
	Add(longUrl string) (string, error)
	Get(short string) (string, error)
}
