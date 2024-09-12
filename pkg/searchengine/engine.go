package searchengine

import (
	"github.com/weissmedia/searchengine/internal/search"
)

// NewEngine stellt eine Instanz der Suchmaschine bereit
func NewEngine(cfg Config) *search.Engine {
	return search.NewEngine(&cfg)
}
