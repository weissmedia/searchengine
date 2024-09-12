package backend

import (
	"fmt"
	"github.com/weissmedia/searchengine/internal/client"
	"github.com/weissmedia/searchengine/internal/core"
	"github.com/weissmedia/searchengine/internal/sorting"
	"golang.org/x/net/context"
	"log"
	"sync"
)

type RedisBackend struct {
	redisClient       *client.RedisClient
	redisSearchClient *client.RedisSearchClient
	sortingMapper     sorting.Mapper
}

func NewRedisBackend(redisClient *client.RedisClient, redisSearchClient *client.RedisSearchClient, sortingMapper sorting.Mapper) *RedisBackend {
	return &RedisBackend{
		redisClient:       redisClient,
		redisSearchClient: redisSearchClient,
		sortingMapper:     sortingMapper,
	}
}

// GetMap returns results for either single or multiple values on a field
func (b *RedisBackend) GetMap(ctx context.Context, field string, value interface{}) (map[string]struct{}, error) {
	switch v := value.(type) {
	case string:
		return b.getMapValue(ctx, field, v)
	case []string:
		return b.getMapValues(ctx, field, v)
	default:
		return nil, fmt.Errorf("invalid value type, expected string or []string")
	}
}

// GetMapExcluding retrieves all field values excluding the given one
func (b *RedisBackend) GetMapExcluding(ctx context.Context, field string, valueExclude interface{}) (map[string]struct{}, error) {
	fieldValues, _, err := b.GetFieldValuesExcluding(ctx, field, valueExclude)
	if err != nil {
		return nil, err
	}

	fields := make([]string, len(fieldValues))
	for i, fieldValue := range fieldValues {
		fields[i] = fmt.Sprintf("%s:%s", field, fieldValue)
	}

	result, err := b.redisClient.SUnionMap(ctx, fields...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetFieldValuesExcluding retrieves field values excluding the given one
func (b *RedisBackend) GetFieldValuesExcluding(ctx context.Context, field string, valueExclude interface{}) ([]string, int, error) {
	// Generate a Redis search pattern to match field values
	pattern := fmt.Sprintf("%s:*", field)
	// Convert the value to exclude into a string
	excludePattern := fmt.Sprintf("%v", valueExclude)

	// Use the redisClient to scan the keys and exclude the given value
	values, total, err := b.redisClient.Scan(ctx, pattern, 4, -1, 0, excludePattern)
	if err != nil {
		return nil, 0, fmt.Errorf("error retrieving field values for %s excluding %v: %v", field, valueExclude, err)
	}

	log.Printf("Found %d field values for %s excluding %v", total, field, valueExclude)
	return values, total, nil
}

// SearchComparisonMap handles a comparison query and executes it using the appropriate operator
func (b *RedisBackend) SearchComparisonMap(field string, operator core.ComparisonOperator, value interface{}) (map[string]struct{}, error) {
	query, err := b.generateComparisonQuery(field, operator, value)
	if err != nil {
		return nil, err
	}

	return b.executeSearch(query)
}

// SearchRangeMap constructs and executes a range query in Redis
func (b *RedisBackend) SearchRangeMap(field string, min, max interface{}) (map[string]struct{}, error) {
	query := fmt.Sprintf("@%s:[%v %v]", field, min, max)
	log.Printf("Generated range query: %s", query)
	return b.executeSearch(query)
}

// SearchFuzzyMap constructs and executes a fuzzy search query in Redis
func (b *RedisBackend) SearchFuzzyMap(field, value string) (map[string]struct{}, error) {
	query := fmt.Sprintf("@%s:%%%%%s%%%%", field, value)
	log.Printf("Generated fuzzy search query: %s", query)
	return b.executeSearch(query)
}

func (b *RedisBackend) SearchWildcardMap(field, value string) (map[string]struct{}, error) {
	// Formatieren des Wildcard-Queries für RedisSearch
	query := fmt.Sprintf("@%s:%s", field, value)
	log.Printf("Executing wildcard search with query: %s", query)
	return b.executeSearch(query)
}

// GetSortedFieldValuesMap retrieves sorted field values using Redis zRange
func (b *RedisBackend) GetSortedFieldValuesMap(ctx context.Context, fields []string) (<-chan core.SortResult, error) {
	zRangeMap, err := b.redisClient.ZRangeMap(ctx, fields)
	if err != nil {
		return nil, err
	}

	resultChan := make(chan core.SortResult, len(fields))
	var wg sync.WaitGroup
	wg.Add(len(fields))

	// Speichere sowohl den Feldnamen als auch den Index
	for i, field := range fields {
		go func(index int, field string) {
			defer wg.Done()

			sortMembers := zRangeMap[field]
			if len(sortMembers) == 0 {
				resultChan <- core.SortResult{
					Field: field, // Behalte den Feldnamen
					Err:   fmt.Errorf("no sorted members for field %s", field),
					Index: index, // Behalte den Index für die Reihenfolge
				}
				return
			}

			orderMap, attrType := b.sortingMapper.CreateSortingMap(sortMembers)
			resultChan <- core.SortResult{
				Index:        index, // Behalte den Index für die Reihenfolge
				Field:        field, // Behalte den Feldnamen
				OrderMap:     orderMap,
				OrderMapType: attrType,
			}
		}(i, field) // Gib sowohl den Index als auch den Feldnamen in die Goroutine
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	return resultChan, nil
}

// Helper to retrieve map values for a single value
func (b *RedisBackend) getMapValue(ctx context.Context, field string, value string) (map[string]struct{}, error) {
	field = fmt.Sprintf("%s:%s", field, value)
	result, err := b.redisClient.SMembersMap(ctx, field)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Helper to retrieve map values for multiple values
func (b *RedisBackend) getMapValues(ctx context.Context, field string, values []string) (map[string]struct{}, error) {
	fields := make([]string, len(values))
	for i, value := range values {
		fields[i] = fmt.Sprintf("%s:%s", field, value)
	}

	result, err := b.redisClient.SUnionMap(ctx, fields...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// generateComparisonQuery constructs the Redis query string for comparison operators like >, >=, <, <=
func (b *RedisBackend) generateComparisonQuery(field string, operator core.ComparisonOperator, value interface{}) (string, error) {
	var query string
	switch operator {
	case core.OpGreater:
		query = fmt.Sprintf("@%s:[(%v +inf]", field, value) // Greater than
	case core.OpGreaterEqual:
		query = fmt.Sprintf("@%s:[%v +inf]", field, value) // Greater than or equal to
	case core.OpLess:
		query = fmt.Sprintf("@%s:[-inf (%v]", field, value) // Less than
	case core.OpLessEqual:
		query = fmt.Sprintf("@%s:[-inf %v]", field, value) // Less than or equal to
	default:
		return "", fmt.Errorf("invalid comparison operator: %s", operator)
	}
	log.Printf("Generated comparison query: %s", query)
	return query, nil
}

// executeSearch performs the actual Redis search and handles pagination if necessary
func (b *RedisBackend) executeSearch(query string) (map[string]struct{}, error) {
	searchResults, total, err := b.redisSearchClient.SearchIDMap(query, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("error executing search: %v", err)
	}

	log.Printf("Initial search completed. Total Results: %v", total)

	if total > 0 {
		searchResults, total, err = b.redisSearchClient.SearchIDMap(query, total, 0)
		if err != nil {
			return nil, fmt.Errorf("error retrieving full results: %v", err)
		}
	}

	log.Printf("Final Search Results: %v", searchResults)
	return searchResults, nil
}
