package search

import (
	"context"
	"github.com/weissmedia/searchengine/generated/sqparser"
	"github.com/weissmedia/searchengine/internal/backend"
	"github.com/weissmedia/searchengine/internal/client"
	"github.com/weissmedia/searchengine/internal/sorting"
	"log"
)

type Engine struct {
	Backend backend.SearchBackend
	parser  *sqparser.SearchQueryParser // ANTLR Parser
}

func NewEngine(cfg backend.Config) *Engine {
	sortingMapper := sorting.NewMapper()
	redisClient := client.NewRedisClient()
	redisSearchClient := client.NewRedisSearchClient(redisClient, cfg.GetIndexName())
	searchBackend := backend.NewRedisBackend(redisClient, redisSearchClient, sortingMapper)
	return &Engine{
		Backend: searchBackend,
	}
}

// Search führt die Suche durch und gibt das ResultSet zurück
func (e *Engine) Search(ctx context.Context, query string) ([]string, error) {
	tree, err := sqparser.Parse(query)
	if err != nil {
		log.Fatal(err)
	}

	exe := NewExecutor(ctx, e.Backend)
	result, err := exe.Execute(tree)
	if err != nil {
		return nil, err
	}
	log.Print("Final Result Visitor: ", result)
	return result, nil
}
