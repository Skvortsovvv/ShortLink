package links

import (
	"fmt"
	"sync"
	"testingTask/internal/shorter"
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
	proceedURL := longURL
	for {
		shortURL = shorter.Shorter(proceedURL)
		repo.mute.RLock()
		origin, ok := repo.data[shortURL]
		repo.mute.RUnlock()
		if !ok {
			repo.mute.Lock()
			repo.data[shortURL] = longURL
			repo.mute.Unlock()
			break
		} else if origin != longURL {
			proceedURL = shortURL
		} else {
			break
		}
	}

	return shortURL, nil
}

func (repo *LinksMemoryRepo) Get(shortULR string) (string, error) {
	repo.mute.RLock()
	longURL, ok := repo.data[shortULR]
	repo.mute.RUnlock()
	if !ok {
		return "", NoURLError
	}
	return longURL, nil
}
