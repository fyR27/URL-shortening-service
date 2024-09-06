package config

import (
	"flag"

	"github.com/google/uuid"
)

type Config struct {
	Host string
	URL  string
}

func NewConfig() *Config {
	return &Config{
		Host: ":8080",
		URL:  uuid.NewString(),
	}
}

func ParseFlags(c *Config) {
	flag.StringVar(&c.Host, "a", ":8080", "address to run server ")
	flag.StringVar(&c.URL, "b", uuid.Newstring(), "url to get base url")
	flag.Parse()
}
