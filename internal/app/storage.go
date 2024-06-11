package app

type Storage struct {
	ID  string
	URL string
}

var m = map[string]string{}

func (s *Storage) AddNewURL(body []byte) string {

	newURL := &Storage{
		ID:  "/EwHXdJfB",
		URL: string(body[:]),
	}

	m[newURL.ID] = newURL.URL
	return newURL.ID
}

func (s *Storage) FindAddr(url string) string {
	if value, ok := m[url]; ok {
		return value
	}
	return url
}
