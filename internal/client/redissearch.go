package client

import (
	"fmt"
	"github.com/RediSearch/redisearch-go/v2/redisearch"
	"github.com/weissmedia/searchengine/internal/core"
	"go.uber.org/zap"
	"strings"
)

type RedisSearchClient struct {
	client       *redisearch.Client
	searchSchema []core.SearchSchema
	prefix       string
	log          *zap.Logger
}

// NewRedisSearchClient creates a new instance of RedisSearchClient
func NewRedisSearchClient(redisClient *RedisClient, searchIndexName string, schema []core.SearchSchema, prefix string, logger *zap.Logger) *RedisSearchClient {
	redis := redisClient.client.Options()
	redisSearchClient := redisearch.NewClient(redis.Addr, searchIndexName)

	r := &RedisSearchClient{
		client:       redisSearchClient,
		searchSchema: schema,
		prefix:       prefix,
		log:          logger,
	}
	r.setRuntimeSetup()

	err := r.createSearchSchema()
	if err != nil {
		r.log.Warn("Error creating search schema", zap.Error(err))
	}

	return r
}

func (r *RedisSearchClient) setRuntimeSetup() {
	setConfig, err := r.client.SetConfig("MAXSEARCHRESULTS", "-1")
	if err != nil {
		r.log.Error("Error setting runtime configuration", zap.Error(err))
	} else {
		r.log.Info("Runtime configuration set", zap.String("config", setConfig))
	}
}

// createSearchSchema uses the searchSchema passed into the client to create the Redis search schema
func (r *RedisSearchClient) createSearchSchema() error {
	return r.setupSchema(r.searchSchema)
}

// setupSchema is a helper to set up the schema, can be reused for both creation and recreation of indexes
func (r *RedisSearchClient) setupSchema(schema []core.SearchSchema) error {
	sc := redisearch.NewSchema(redisearch.DefaultOptions)
	for _, field := range schema {
		switch field.Type {
		case core.TextField:
			sc = sc.AddField(redisearch.NewTextFieldOptions(field.Path, redisearch.TextFieldOptions{As: field.Name}))
		case core.NumericField:
			sc = sc.AddField(redisearch.NewNumericFieldOptions(field.Path, redisearch.NumericFieldOptions{As: field.Name}))
		}
	}

	// Define the index with JSON support
	definition := redisearch.IndexDefinition{IndexOn: "JSON", Prefix: []string{r.getPrefix()}}
	if err := r.client.CreateIndexWithIndexDefinition(sc, &definition); err != nil {
		return fmt.Errorf("error creating index: %v", err)
	}
	r.log.Info("Search schema created", zap.String("prefix", r.getPrefix()))
	return nil
}

func (r *RedisSearchClient) getPrefix() string {
	return r.prefix // Use the configurable prefix
}

// RecreateRedisearchIndex handles the recreation of the RediSearch index based on the provided search schema.
func (r *RedisSearchClient) RecreateRedisearchIndex(indexName string) error {
	// Check if the index exists before attempting to drop it
	indices, err := r.client.List()
	if err != nil {
		return fmt.Errorf("error listing indices: %v", err)
	}

	indexExists := false
	for _, index := range indices {
		if index == indexName {
			indexExists = true
			break
		}
	}

	// If the index exists, attempt to drop it
	if indexExists {
		if err := r.client.DropIndex(false); err != nil {
			return fmt.Errorf("error dropping existing index: %v", err)
		}
		r.log.Info("Existing index dropped", zap.String("index", indexName))
	} else {
		r.log.Info("Index does not exist, skipping drop", zap.String("index", indexName))
	}

	// Recreate the schema and index
	if err := r.setupSchema(r.searchSchema); err != nil {
		return fmt.Errorf("error recreating index: %v", err)
	}

	r.log.Info("Index recreated successfully", zap.String("index", indexName))
	return nil
}

func (r *RedisSearchClient) SearchIDMap(query string, limit, offset int) (map[string]struct{}, int, error) {
	q := redisearch.NewQuery(query)
	q.SetFlags(redisearch.QueryNoContent)
	q.Limit(offset, limit)
	docs, total, err := r.client.Search(q)
	if err != nil {
		r.log.Error("Error during search", zap.Error(err))
		return nil, 0, err
	}
	r.log.Info("Search results", zap.Any("docs", docs), zap.Int("total", total))

	// Prefix to be removed from result Ids
	prefix := r.getPrefix()

	resultSet := make(map[string]struct{}, total)
	for _, result := range docs {
		// Remove the prefix from the result id
		cleanedID := strings.TrimPrefix(result.Id, prefix+":")
		resultSet[cleanedID] = struct{}{}
	}
	r.log.Debug("Search Results", zap.Any("resultSet", resultSet))
	return resultSet, total, nil
}
