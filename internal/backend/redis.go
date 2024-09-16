package backend

import (
	"fmt"
	"github.com/weissmedia/searchengine/internal/client"
	"github.com/weissmedia/searchengine/internal/core"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"sync"
)

type RedisBackend struct {
	redisClient       *client.RedisClient
	redisSearchClient *client.RedisSearchClient
	log               *zap.Logger
}

// NewRedisBackend initializes a new RedisBackend instance with Redis clients, sorting mapper, and logger.
func NewRedisBackend(redisClient *client.RedisClient, redisSearchClient *client.RedisSearchClient, logger *zap.Logger) *RedisBackend {
	return &RedisBackend{
		redisClient:       redisClient,
		redisSearchClient: redisSearchClient,
		log:               logger,
	}
}

// GetMap returns the results for a specific field based on a value, supporting both single string and string array values.
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

// GetMapExcluding retrieves all field values except the one specified.
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

// GetFieldValuesExcluding retrieves field values excluding the specified value.
func (b *RedisBackend) GetFieldValuesExcluding(ctx context.Context, field string, valueExclude interface{}) ([]string, int, error) {
	pattern := fmt.Sprintf("%s:*", field)
	excludePattern := fmt.Sprintf("%v", valueExclude)

	// Scans the keys while excluding the specified value
	values, total, err := b.redisClient.Scan(ctx, pattern, 4, -1, 0, excludePattern)
	if err != nil {
		return nil, 0, fmt.Errorf("error retrieving field values for %s excluding %v: %v", field, valueExclude, err)
	}

	b.log.Info("Found field values", zap.Int("total", total), zap.String("field", field), zap.Any("excludedValue", valueExclude))
	return values, total, nil
}

// SearchComparisonMap constructs and executes a comparison query for the specified field using the provided operator.
func (b *RedisBackend) SearchComparisonMap(field string, operator core.ComparisonOperator, value interface{}) (map[string]struct{}, error) {
	query, err := b.generateComparisonQuery(field, operator, value)
	if err != nil {
		return nil, err
	}

	return b.executeSearch(query)
}

// SearchRangeMap constructs and executes a range query for the specified field between a minimum and maximum value.
func (b *RedisBackend) SearchRangeMap(field string, min, max interface{}) (map[string]struct{}, error) {
	query := fmt.Sprintf("@%s:[%v %v]", field, min, max)
	b.log.Info("Generated range query", zap.String("query", query))
	return b.executeSearch(query)
}

// SearchFuzzyMap constructs and executes a fuzzy search query for the specified field.
func (b *RedisBackend) SearchFuzzyMap(field, value string) (map[string]struct{}, error) {
	query := fmt.Sprintf("@%s:%%%%%s%%%%", field, value)
	b.log.Info("Generated fuzzy search query", zap.String("query", query))
	return b.executeSearch(query)
}

// SearchWildcardMap constructs and executes a wildcard search query for the specified field.
func (b *RedisBackend) SearchWildcardMap(field, value string) (map[string]struct{}, error) {
	query := fmt.Sprintf("@%s:%s", field, value)
	b.log.Info("Executing wildcard search", zap.String("query", query))
	return b.executeSearch(query)
}

// GetSortedFieldValuesMap retrieves sorted field values using Redis ZRANGE and returns them via a channel.
func (b *RedisBackend) GetSortedFieldValuesMap(ctx context.Context, sortFields *core.SortFieldList) (<-chan core.SortResult, error) {
	fields := sortFields.SortFields()
	zRangeMap, err := b.redisClient.ZRangeMap(ctx, fields)
	if err != nil {
		return nil, err
	}

	resultChan := make(chan core.SortResult, len(fields))
	var wg sync.WaitGroup
	wg.Add(len(fields))

	// Launch goroutines to process each field
	for i, field := range fields {
		go func(index int, field string) {
			defer wg.Done()

			sortMembers := zRangeMap[field]
			if len(sortMembers) == 0 {
				resultChan <- core.SortResult{
					Field: field,
					Err:   fmt.Errorf("no sorted members for field %s", field),
					Index: index,
				}
				return
			}

			orderMap, attrType := sortFields.BuildSortMap(sortMembers)
			resultChan <- core.SortResult{
				Index:        index,
				Field:        field,
				OrderMap:     orderMap,
				OrderMapType: attrType,
			}
		}(i, field)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	return resultChan, nil
}

// getMapValue is a helper method to retrieve values for a single field.
func (b *RedisBackend) getMapValue(ctx context.Context, field string, value string) (map[string]struct{}, error) {
	field = fmt.Sprintf("%s:%s", field, value)
	result, err := b.redisClient.SMembersMap(ctx, field)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getMapValues is a helper method to retrieve values for multiple fields.
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

// generateComparisonQuery generates the Redis search query for comparison operations like >, >=, <, <=.
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
	b.log.Info("Generated comparison query", zap.String("query", query))
	return query, nil
}

// executeSearch performs a Redis search query and retrieves the results.
func (b *RedisBackend) executeSearch(query string) (map[string]struct{}, error) {
	searchResults, total, err := b.redisSearchClient.SearchIDMap(query, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("error executing search: %v", err)
	}

	b.log.Info("Initial search completed", zap.Int("totalResults", total))

	if total > 0 {
		searchResults, total, err = b.redisSearchClient.SearchIDMap(query, total, 0)
		if err != nil {
			return nil, fmt.Errorf("error retrieving full results: %v", err)
		}
	}

	b.log.Info("Final search results", zap.Any("results", searchResults))
	return searchResults, nil
}

// UpdateSearchIndex updates the Redis search index by recreating it.
func (b *RedisBackend) UpdateSearchIndex(indexName string) (bool, error) {
	err := b.redisSearchClient.RecreateRedisearchIndex(indexName)
	if err != nil {
		b.log.Error("Failed to recreate index", zap.String("indexName", indexName), zap.Error(err))
		return false, err
	}
	return true, nil
}
