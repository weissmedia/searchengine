package search

import (
	"context"
	"github.com/weissmedia/searchengine/generated/sqparser"
	"github.com/weissmedia/searchengine/internal/backend"
	"github.com/weissmedia/searchengine/internal/client"
	"github.com/weissmedia/searchengine/internal/core"
	"github.com/weissmedia/searchengine/internal/profiler"
	"go.uber.org/zap"
)

type Engine struct {
	Backend  backend.SearchBackend
	parser   *sqparser.SearchQueryParser // ANTLR Parser
	log      *zap.Logger
	profiler *profiler.Profiler
}

// NewEngine creates a new search engine instance
func NewEngine(cfg backend.Config, logger *zap.Logger) *Engine {
	logger.Info("Creating new search engine instance...")
	prof := profiler.NewProfiler(cfg.GetProfilerEnabled(), logger)
	redisClient := client.NewRedisClient(
		cfg.GetRedisHost(),     // Redis host from config
		cfg.GetRedisPort(),     // Redis port from config
		cfg.GetRedisDB(),       // Redis DB from config
		cfg.GetFilterPrefix(),  // Filter prefix from config
		cfg.GetSortingPrefix(), // Sorting prefix from config
		logger,
	)

	redisSearchClient := client.NewRedisSearchClient(
		redisClient,
		cfg.GetSearchIndexName(), // Search index from config
		cfg.GetSearchSchema(),    // Search schema from config
		cfg.GetDataPrefix(),      // Data prefix from config
		logger,
	)

	searchBackend := backend.NewRedisBackend(redisClient, redisSearchClient, logger)

	return &Engine{
		Backend:  searchBackend,
		log:      logger,
		profiler: prof,
	}
}

// Search executes the search and returns the result set
func (e *Engine) Search(ctx context.Context, query string) (*core.ExecutionResult, error) {
	// Parse the query
	tree, err := sqparser.Parse(query)
	if err != nil {
		e.log.Error("Error parsing query", zap.String("query", query), zap.Error(err))
		return nil, err // Return early if there's a parsing error
	}

	// Execute the query
	exe := NewExecutor(ctx, e.Backend, e.log, e.profiler)
	result, err := exe.Execute(tree)
	if err != nil {
		e.log.Error("Error executing query", zap.String("query", query), zap.Error(err))
		return nil, err
	}

	// Log the final result and return
	e.log.Info("Query executed successfully", zap.String("query", query), zap.Any("result", result))
	return result, nil
}
