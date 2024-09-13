package searchengine

import (
	cfg "github.com/weissmedia/searchengine/internal/config"
)

// NewConfig creates a new instance of Config by reading environment variables and loading the search schema file.
func NewConfig() (*cfg.Config, error) {
	return cfg.NewConfig()
}
