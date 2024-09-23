package app

import (
	"github.com/fyR27/URL-shortening-service/config"
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

func ValidURL(c *config.Config) string {
	if c.URL[len(c.URL)-1] >= 48 && c.URL[len(c.URL)-1] <= 57 {
		return c.URL + "/"
	} else {
		return c.URL + c.Host + "/"
	}
}

func (s *Storage) AddNewURL(body []byte, c *config.Config) string {
	newURL := &ParsedURL{
		ID:  uuid.NewString(),
		URL: string(body[:]),
	}
	s.store[newURL.ID] = newURL.URL
	return ValidURL(c) + newURL.ID
}

func (s *Storage) FindAddr(url string) string {
	if value, ok := s.store[url]; ok {
		return value
	}
	return "Bad id"
}
