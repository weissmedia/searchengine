package query

import (
	"github.com/weissmedia/searchengine/generated/sqparser"
	"github.com/weissmedia/searchengine/internal/backend"
	"github.com/weissmedia/searchengine/searchengine/config"
	"log"
)

type SearchEngine struct {
	Backend backend.SearchBackend
	parser  *sqparser.SearchQueryParser // ANTLR Parser
}

func NewSearchEngine(cfg config.Config) *SearchEngine {
	searchBackend := backend.NewRedisBackend(cfg.RedisAddress)
	return &SearchEngine{
		Backend: searchBackend,
	}
}

// Search führt die Suche durch und gibt das ResultSet zurück
func (se *SearchEngine) Search(query string) ([]string, error) {
	// Beispiel-Query
	query = "keyA IN ('val1', 'val2', 'val3')"
	log.Println("Query:", query)

	tree, err := sqparser.Parse(query)
	if err != nil {
		log.Fatal(err)
	}

	exe := NewExecutor(se.Backend)
	result, err := exe.Execute(tree)
	if err != nil {
		return nil, err
	}
	log.Print("Final Result Visitor: ", result)
	return result, nil
}
