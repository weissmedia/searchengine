package main

import (
	"fmt"
	"github.com/weissmedia/searchengine/pkg/searchengine"
	"golang.org/x/net/context"
	"log"
	"os"
)

func main() {
	os.Setenv("SEARCH_SCHEMA_FILE", "./example/searchschema.json")
	os.Setenv("SEARCH_INDEX_NAME", "idx2")
	os.Setenv("NAMESPACE_PREFIX", "opus.bdl.datapool")
	os.Setenv("REDIS_FILTER_PREFIX", "opus.bdl.datapool:filter:")
	os.Setenv("REDIS_SORTING_PREFIX", "opus.bdl.datapool:sorting:")

	// Load the configuration, including the schema from the file
	cfg, err := searchengine.NewConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v\n", err)
	}

	// Use the configuration (e.g., connect to Redis, use schema)
	fmt.Printf("Loaded Config: %+v\n", cfg)

	engine := searchengine.NewEngine(cfg)
	ctx := context.Background()
	search, err := engine.Search(ctx, "inbox = 'AR_IB_AK_VL_KB_AHV'")
	if err != nil {
		log.Fatalf("Error query: %v\n", err)
	}
	fmt.Println(search)
}
