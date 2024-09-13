package search

import (
	"context"
	"github.com/weissmedia/searchengine/generated/sqparser"
	"github.com/weissmedia/searchengine/internal/backend"
	"github.com/weissmedia/searchengine/internal/client"
	"github.com/weissmedia/searchengine/internal/log"
	"github.com/weissmedia/searchengine/internal/sorting"
	"go.uber.org/zap"
)

type Engine struct {
	Backend backend.SearchBackend
	parser  *sqparser.SearchQueryParser // ANTLR Parser
	log     *zap.Logger
}

func NewEngine(cfg backend.Config) *Engine {
	logger := log.GetLogger()
	logger.Info("Creating new search engine instance...")
	sortingMapper := sorting.NewMapper()
	redisClient := client.NewRedisClient(cfg.GetRedisHost(), cfg.GetRedisPort(), cfg.GetRedisDB(), cfg.GetFilterPrefix(), cfg.GetSortingPrefix())
	redisSearchClient := client.NewRedisSearchClient(redisClient, cfg.GetSearchIndexName(), cfg.GetSearchSchema(), cfg.GetDataPrefix())
	searchBackend := backend.NewRedisBackend(redisClient, redisSearchClient, sortingMapper)
	return &Engine{
		Backend: searchBackend,
		log:     logger,
	}
}

// Search führt die Suche durch und gibt das ResultSet zurück
func (e *Engine) Search(ctx context.Context, query string) ([]string, error) {
	tree, err := sqparser.Parse(query)
	if err != nil {
		e.log.Error("Error parsing query", zap.Error(err))
	}

	exe := NewExecutor(ctx, e.Backend)
	result, err := exe.Execute(tree)
	if err != nil {
		e.log.Error("Error executing query", zap.Error(err))
		return nil, err
	}
	e.log.Info("Final Result Visitor", zap.Any("result", result))
	return result, nil
}
