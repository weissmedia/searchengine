package searchengine

import (
	"github.com/weissmedia/searchengine/internal/backend"
	"github.com/weissmedia/searchengine/internal/search"
	"go.uber.org/zap"
)

// NewEngine initializes and returns an instance of the search engine.
// It takes the configuration and logger as parameters to configure and log search engine events.
func NewEngine(cfg backend.Config, logger *zap.Logger) *search.Engine {
	return search.NewEngine(cfg, logger)
}
