package pkg

import (
	"github.com/weissmedia/searchengine/internal/query"
	"github.com/weissmedia/searchengine/pkg/config"
)

// NewSearchEngine stellt eine Instanz der Suchmaschine bereit
func NewSearchEngine(cfg config.Config) *query.SearchEngine {
	return query.NewSearchEngine(cfg)
}
