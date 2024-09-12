package client

import (
	"fmt"
	"github.com/RediSearch/redisearch-go/v2/redisearch"
	"github.com/weissmedia/searchengine/internal/core"
	"log"
	"strings"
)

type RedisSearchClient struct {
	client       *redisearch.Client
	searchSchema []core.SearchSchema
}

// NewRedisSearchClient erstellt eine neue Instanz des RedisSearchClient
func NewRedisSearchClient(redisClient *RedisClient, searchIndexName string) *RedisSearchClient {
	redis := redisClient.client.Options()
	redisSearchClient := redisearch.NewClient(redis.Addr, searchIndexName)

	r := &RedisSearchClient{client: redisSearchClient}
	r.setRuntimeSetup()
	r.setSearchSchemaAttributes()

	err := r.createSearchSchema()
	if err != nil {
		log.Println(err)
	}

	return r
}

func (r *RedisSearchClient) setRuntimeSetup() {
	setConfig, err := r.client.SetConfig("MAXSEARCHRESULTS", "-1")
	if err != nil {
		log.Println(err)
	} else {
		log.Println(setConfig)
	}
}

func (r *RedisSearchClient) setSearchSchemaAttributes() {
	r.searchSchema = []core.SearchSchema{
		attribute("id", core.TextField),
		attribute("keyA", core.TextField),
		attribute("keyB", core.TextField),
		attribute("keyC", core.TextField),
		attribute("keyD", core.TextField),
		attribute("keyF", core.TextField),
		attribute("keyE", core.NumericField),
		attribute("num", core.NumericField),
	}
}

// NewFieldInfo creates a new FieldInfo instance
func attribute(path string, fieldType core.SearchSchemaType) core.SearchSchema {
	var searchOptions string
	switch fieldType {
	case core.TextField:
		searchOptions = core.TextSearchOptions
	case core.NumericField:
		searchOptions = core.NumericSearchOptions
	default:
		searchOptions = ""
	}
	return core.SearchSchema{Path: fmt.Sprintf("$.%s", path), Type: fieldType, SearchOptions: searchOptions, Name: strings.ReplaceAll(path, ".", "_")}
}

func (r *RedisSearchClient) createSearchSchema() error {
	sc := redisearch.NewSchema(redisearch.DefaultOptions)
	for _, field := range r.searchSchema {
		switch field.Type {
		case core.TextField:
			sc = sc.AddField(redisearch.NewTextFieldOptions(field.Path, redisearch.TextFieldOptions{As: field.Name}))
		case core.NumericField:
			sc = sc.AddField(redisearch.NewNumericFieldOptions(field.Path, redisearch.NumericFieldOptions{As: field.Name}))
		}
	}
	// ToDo: Make the prefix configurable as quickly as possible
	definition := redisearch.IndexDefinition{IndexOn: "JSON", Prefix: []string{r.getPrefix()}}
	if err := r.client.CreateIndexWithIndexDefinition(sc, &definition); err != nil {
		return fmt.Errorf("error creating index: %v", err)
	}
	return nil
}

func (r *RedisSearchClient) getPrefix() string {
	// todo: make it configurable as quickly as possible
	return fmt.Sprint("opus.bdl.datapool:data:")
}

func (r *RedisSearchClient) GetSearchAttributes() []core.SearchSchema {
	return r.searchSchema
}

func (r *RedisSearchClient) SearchIDMap(query string, limit, offset int) (map[string]struct{}, int, error) {
	q := redisearch.NewQuery(query)
	q.SetFlags(redisearch.QueryNoContent)
	q.Limit(offset, limit)
	docs, total, err := r.client.Search(q)
	if err != nil {
		return nil, 0, err
	}
	log.Printf("search results: %v %v %v", docs, total, err)
	// Prefix to be removed from result Ids
	prefix := r.getPrefix()

	resultSet := make(map[string]struct{}, total)
	for _, result := range docs {
		// Remove the prefix from the result.Id
		cleanedID := strings.TrimPrefix(result.Id, prefix)
		resultSet[cleanedID] = struct{}{}
	}
	log.Printf("Search Results: %v\n", resultSet)
	return resultSet, total, nil
}
