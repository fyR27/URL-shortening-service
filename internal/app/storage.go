package app

import (
	"fmt"

	"github.com/google/uuid"
)

type ParserURL struct {
	ID  string
	URL string
}

var storage = map[string]string{}

func (p *ParserURL) AddNewURL(body []byte) string {
	newURL := &ParserURL{
		ID:  uuid.NewString(),
		URL: string(body[:]),
	}
	storage[newURL.ID] = newURL.URL
	fmt.Println(uuid.SetNodeID([]byte(newURL.URL)))
	return newURL.ID
}

func (p *ParserURL) FindAddr(url string) string {
	if value, ok := storage[url]; ok {
		return value
	}
	return url
}
