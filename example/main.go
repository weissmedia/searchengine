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

	query := "data_keyE=[0 100]"
	fmt.Println("Query:", query)
	// Execute a search query
	searchResult, err := engine.Search(ctx, query)
	if err != nil {
		logger.Error("Error executing query", zap.Error(err))
		return
	}

	// Output the search result
	fmt.Println("Search Results:", searchResult.ResultSet)

	// Schöner formatierte Ausgabe der Timings
	fmt.Println("Operation Timings:")
	for _, timing := range searchResult.Timings {
		fmt.Printf("  Operation: %s | Time: %.3f ms\n", timing.Operation, timing.TimeMs)
	}

	// Ausgabe der Gesamtzeit
	fmt.Printf("Total Time: %.3f ms\n", searchResult.TotalTime)

	// Ensure all log entries are written before exiting
	log.SyncLogger()
}
