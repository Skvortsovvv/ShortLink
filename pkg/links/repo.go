package links

import (
	"fmt"
	"net/http"
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

func (repo *LinksMemoryRepo) Add(longURL string, r *http.Request) (string, error) {
	// generated shortURL with method
	var shortURL string
	repo.mute.RLock()
	_, ok := repo.data[shortURL]
	repo.mute.Unlock()
	if !ok {
		repo.mute.Lock()
		repo.data[shortURL] = longURL
		repo.mute.Unlock()
	}

	return shortURL, nil
}

func (repo *LinksMemoryRepo) Get(shortULR string, r *http.Request) (string, error) {
	repo.mute.RLock()
	_, ok := repo.data[shortULR]
	repo.mute.Unlock()
	if !ok {
		return "", NoURLError
	}
	repo.mute.RLock()
	longURL := repo.data[shortULR]
	repo.mute.RUnlock()
	return longURL, nil
}
