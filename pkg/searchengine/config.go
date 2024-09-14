package searchengine

import (
	"github.com/weissmedia/searchengine/internal/backend"
	cfg "github.com/weissmedia/searchengine/internal/config"
)

// NewConfig creates a new instance of Config by reading environment variables and loading the search schema file.
func NewConfig() (backend.Config, error) {
	return cfg.NewConfig()
}
