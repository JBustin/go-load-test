package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var DefaultConfig Config = Config{
	IsBrowser:   false,
	IsSerie:     false,
	Hits:        100,
	WaitMs:      1000,
	Concurrency: 50,
	TimeoutMs:   20000,
	Scrap:       false,
	Urls:        []string{},
	Headers:     make(map[string]string),
	LogLevel:    "error",
}

type Config struct {
	IsBrowser   bool
	IsSerie     bool
	Hits        int
	WaitMs      int
	Concurrency int
	TimeoutMs   int
	Scrap       bool
	Urls        []string
	Headers     map[string]string
	LogLevel    string
}

func (c Config) String() string {
	return fmt.Sprintf(`
	IsBrowser: %v
	IsSerie: %v
	Hits: %v
	WaitMs: %v
	Concurrency: %v
	TimeoutMs: %v
	Scrap: %v
	Urls: %v
	Headers: %v
	LogLevel: %v
	`, c.IsBrowser,
		c.IsSerie,
		c.Hits,
		c.WaitMs,
		c.Concurrency,
		c.TimeoutMs,
		c.Scrap,
		c.Urls,
		c.Headers,
		c.LogLevel)
}

func New(jsonFilePath string) (Config, error) {
	var c Config

	content, err := ioutil.ReadFile(jsonFilePath)

	if err != nil {
		return c, err
	}

	c = DefaultConfig
	err = json.Unmarshal(content, &c)

	return c, err
}
