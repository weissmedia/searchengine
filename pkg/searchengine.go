package pkg

import (
	"github.com/weissmedia/searchengine/internal/config"
	"github.com/weissmedia/searchengine/internal/query"
)

// NewSearchEngine stellt eine Instanz der Suchmaschine bereit
func NewSearchEngine(cfg config.Config) *query.SearchEngine {
	return query.NewSearchEngine(cfg)
}
