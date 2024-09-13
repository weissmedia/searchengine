package searchengine

import (
	cfg "github.com/weissmedia/searchengine/internal/config"
	"github.com/weissmedia/searchengine/internal/search"
)

// NewEngine stellt eine Instanz der Suchmaschine bereit
func NewEngine(cfg *cfg.Config) *search.Engine {
	return search.NewEngine(cfg)
}
