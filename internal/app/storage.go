package app

import (
	"regexp"

	"github.com/fyR27/URL-shortening-service/config"
	"github.com/google/uuid"
)

var re = regexp.MustCompile(`https?:\/\/([a-zA-Z0-9-]+)(\.[a-zA-Z0-9-]+)*(:[0-9])?`)

type ParsedURL struct {
	ID  string
	URL string
}

type Storage struct {
	store   map[string]string
	baseURL string
}

func validURL(c *config.Config) string {
	matched := re.MatchString(c.URL)
	if matched == bool(true) {
		return c.URL + "/"
	}
	return c.URL + c.Host + "/"
}

func NewStore(c *config.Config) *Storage {
	return &Storage{
		store:   make(map[string]string),
		baseURL: validURL(c),
	}
}

func (s *Storage) AddNewURL(body []byte) string {
	newURL := &ParsedURL{
		ID:  uuid.NewString(),
		URL: string(body[:]),
	}
	s.store[newURL.ID] = newURL.URL
	return s.baseURL + newURL.ID
}

func (s *Storage) FindAddr(url string) string {
	if value, ok := s.store[url]; ok {
		return value
	}
	return "Bad id"
}
