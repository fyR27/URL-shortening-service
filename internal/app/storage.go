package app

import (
	"regexp"

	"github.com/fyR27/URL-shortening-service/config"
	"github.com/google/uuid"
)

type ParsedURL struct {
	ID  string
	URL string
}

type Storage struct {
	store   map[string]string
	baseURl string
}

func validURL(c *config.Config) string {
	_, err := regexp.MatchString("[1-9]", c.URL[len(c.URL)-4:len(c.URL)])
	if err != nil {
		return c.URL + "/"
	}
	return c.URL + c.Host + "/"
}

func NewStore(c *config.Config) *Storage {
	return &Storage{
		store:   make(map[string]string),
		baseURl: validURL(c),
	}
}

func (s *Storage) AddNewURL(body []byte) string {
	newURL := &ParsedURL{
		ID:  uuid.NewString(),
		URL: string(body[:]),
	}
	s.store[newURL.ID] = newURL.URL
	return s.baseURl + newURL.ID
}

func (s *Storage) FindAddr(url string) string {
	if value, ok := s.store[url]; ok {
		return value
	}
	return "Bad id"
}
