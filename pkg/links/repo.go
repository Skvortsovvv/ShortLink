package links

import (
	"fmt"
	"sync"
)

var (
	NoURLError = fmt.Errorf("original url not found")
)

type LinksMemoryRepo struct {
	data map[string]string
	mute *sync.RWMutex
}

func NewLinksMemoryRepo() *LinksMemoryRepo {
	return &LinksMemoryRepo{
		data: make(map[string]string),
		mute: &sync.RWMutex{},
	}
}

func (repo *LinksMemoryRepo) Add(longURL string) (string, error) {

	var shortURL string
	proceed := longURL
	for {
		shortURL = Shorter(proceed)
		repo.mute.RLock()
		origin, ok := repo.data[shortURL]
		repo.mute.Unlock()
		if !ok {
			repo.mute.Lock()
			repo.data[shortURL] = longURL
			repo.mute.Unlock()
			break
		} else if origin != longURL {
			proceed = shortURL
		} else {
			break
		}
	}

	return shortURL, nil
}

func (repo *LinksMemoryRepo) Get(shortULR string) (string, error) {
	repo.mute.RLock()
	longURL, ok := repo.data[shortULR]
	repo.mute.Unlock()
	if !ok {
		return "", NoURLError
	}
	return longURL, nil
}
