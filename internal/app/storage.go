package app

import (
	"github.com/google/uuid"
)

type ParsedURL struct {
	ID  string
	URL string
}

type Storage struct {
	store map[string]string
}

func NewStore() *Storage {
	return &Storage{
		store: make(map[string]string),
	}
}

func (s *Storage) AddNewURL(body []byte) string {
	newURL := &ParsedURL{
		ID:  uuid.NewString(),
		URL: string(body[:]),
	}
	s.store[newURL.ID] = newURL.URL
	return newURL.ID
}

func (s *Storage) FindAddr(url string) string {
	if value, ok := s.store[url]; ok {
		return value
	}
	return "Bad id"
}
