package main

import (
	"fmt"
	"github.com/weissmedia/searchengine/internal/log"
	"github.com/weissmedia/searchengine/pkg/searchengine"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

func main() {
	// Retrieve the globally initialized logger
	logger := log.GetLogger()

	// Load the configuration
	cfg, err := searchengine.NewConfig()
	if err != nil {
		logger.Error("Error loading config", zap.Error(err))
		return
	}

	// Initialize the search engine with the configuration and logger
	engine := searchengine.NewEngine(cfg, logger)
	ctx := context.Background()

	// Execute a search query
	searchResult, err := engine.Search(ctx, "data_keyE=[0 100]")
	if err != nil {
		logger.Error("Error executing query", zap.Error(err))
		return
	}

	// Output the search result
	fmt.Println(searchResult.ResultSet)

	// Ensure all log entries are written before exiting
	defer log.SyncLogger()
}
