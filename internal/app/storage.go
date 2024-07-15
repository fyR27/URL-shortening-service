package app

import (
	"fmt"

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
func (s *Storage) AddNewURL(url []byte) string {
	newURL := &ParsedURL{
		ID:  uuid.NewString(),
		URL: string(url),
	}
	s.store[newURL.ID] = newURL.URL
	fmt.Println(s.store)
	return newURL.ID
}

func (s *Storage) FindAddr(id string) string {

	if value, ok := s.store[id]; ok {
		return value
	}
	return "Bad id"
}
