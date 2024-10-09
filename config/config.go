package config

import (
	"flag"
)

type Config struct {
	Host string
	URL  string
}

func NewConfig() *Config {
	return &Config{
		Host: ":8080",
		URL:  "http://localhost",
	}
}

func ParseFlags(c *Config) {
	flag.StringVar(&c.Host, "a", ":8080", "address to run server ")
	flag.StringVar(&c.URL, "b", "http://localhost", "url to get base url")
	flag.Parse()
}
