package app

type Storage struct {
	ID  string
	URL string
}

var m = map[string]string{}

func (s *Storage) AddNewUrl(body []byte) string {

	newUrl := &Storage{
		ID:  "/EwHXdJfB",
		URL: string(body[:]),
	}

	m[newUrl.ID] = newUrl.URL
	return newUrl.ID
}

func (s *Storage) FindAddr(url string) string {
	var key string
	for i, v := range m {
		if i == url {
			key = v
		}
	}
	return key
}
