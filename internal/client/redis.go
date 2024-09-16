package client

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"sort"
	"strings"
)

type RedisClient struct {
	client        *redis.Client
	fieldMappings map[string]string
	prefixFilter  string
	prefixSorting string
	log           *zap.Logger
}

// NewRedisClient creates a new Redis client using the given host, port, and Redis DB.
func NewRedisClient(host string, port, db int, prefixFilter, prefixSorting string, logger *zap.Logger) *RedisClient {
	address := fmt.Sprintf("%s:%d", host, port) // Construct the address from host and port

	client := redis.NewClient(&redis.Options{
		Addr: address, // Use constructed address
		DB:   db,      // Use the provided Redis DB
	})

	return &RedisClient{
		client:        client,
		prefixFilter:  prefixFilter,
		prefixSorting: prefixSorting,
		log:           logger,
	}
}

// SMembersMap retrieves all members of a Redis set and returns them as a map.
func (c *RedisClient) SMembersMap(ctx context.Context, key string) (map[string]struct{}, error) {
	res := c.client.SMembersMap(ctx, c.createFilterField(key))
	return res.Result()
}

// Scan performs a scan operation over Redis keys that match the given pattern, filtering out specific values and supporting offset and limit for pagination.
func (c *RedisClient) Scan(ctx context.Context, pattern string, patternIndex int, limit int, offset int, excludeValue string) ([]string, int, error) {
	pattern = c.createFilterField(pattern)
	var cursor uint64 = 0

	// Slice to store field values
	fieldValues := make([]string, 0)

	// Scan keys and collect attributes
	iter := c.client.Scan(ctx, cursor, pattern, 0).Iterator()
	for iter.Next(ctx) {
		key := iter.Val()
		parts := strings.Split(key, ":")

		// Check if the value should be excluded
		if len(parts) >= patternIndex {
			attr := parts[patternIndex-1]

			// Exclude only if a valid `excludeValue` is set
			if excludeValue == "" || attr != excludeValue {
				fieldValues = append(fieldValues, attr)
			}
		} else {
			c.log.Warn("Pattern index out of range for key", zap.Int("patternIndex", patternIndex), zap.String("key", key))
		}
	}

	if err := iter.Err(); err != nil {
		c.log.Error("Error iterating over filter fields", zap.Error(err))
		return nil, 0, err
	}

	// Sort the collected attributes
	sort.Strings(fieldValues)

	// Calculate total count of unique attributes
	totalCount := len(fieldValues)

	// If the limit is -1, fetch all remaining items after offset
	if limit == -1 {
		limit = totalCount - offset
	}

	// Apply limit and offset to attributes
	start := offset
	if start > totalCount {
		start = totalCount
	}
	end := start + limit
	if end > totalCount {
		end = totalCount
	}
	limitedAttributes := fieldValues[start:end]

	return limitedAttributes, totalCount, nil
}

// SUnionMap performs the Redis SUNION command on multiple sets and returns the union as a map.
func (c *RedisClient) SUnionMap(ctx context.Context, keys ...string) (map[string]struct{}, error) {
	// Create a copy of the keys
	keysCopy := make([]string, len(keys))
	copy(keysCopy, keys)

	// Update the keys in the copy
	for i, key := range keysCopy {
		keysCopy[i] = c.createFilterField(key) // Apply filter field for each key
	}

	// Execute the SUNION command with the copied keys
	result, err := c.client.SUnion(ctx, keysCopy...).Result()
	if err != nil {
		c.log.Error("Error executing SUNION", zap.Error(err))
		return nil, err
	}

	// Convert the result (a list of strings) into a map
	unionMap := make(map[string]struct{}, len(result))
	for _, item := range result {
		unionMap[item] = struct{}{}
	}

	return unionMap, nil
}

// ZRangeIndexMap retrieves and returns the sorted index map for multiple keys using the Redis ZRANGE command.
func (c *RedisClient) ZRangeIndexMap(ctx context.Context, keys []string) ([]map[string]int, error) {
	pipe := c.client.Pipeline()
	sortCmds := make([]*redis.StringSliceCmd, len(keys))
	for i, key := range keys {
		sortCmds[i] = pipe.ZRange(ctx, c.createSortField(key), 0, -1)
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		c.log.Error("Error executing ZRange pipeline", zap.Error(err))
		return nil, err
	}

	sortedIndexMaps := make([]map[string]int, len(keys))

	for i, key := range keys {
		members, err := sortCmds[i].Result()
		if err != nil {
			c.log.Error("Error fetching sorted members", zap.String("key", key), zap.Error(err))
			return nil, fmt.Errorf("error fetching sorted members for key %s: %w", key, err)
		}
		sortedIndexMap := make(map[string]int)
		for index, value := range members {
			val := strings.Split(value, ":")[1]
			sortedIndexMap[val] = index
		}
		sortedIndexMaps[i] = sortedIndexMap
	}

	return sortedIndexMaps, nil
}

// ZRangeMap retrieves the sorted sets for multiple keys using the Redis ZRANGE command and returns them as a map.
func (c *RedisClient) ZRangeMap(ctx context.Context, keys []string) (map[string][]string, error) {
	pipe := c.client.Pipeline()
	sortCmds := make([]*redis.StringSliceCmd, len(keys))

	for i, key := range keys {
		sortCmds[i] = pipe.ZRange(ctx, c.createSortField(key), 0, -1)
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		c.log.Error("Error executing pipeline", zap.Error(err))
		return nil, fmt.Errorf("error executing pipeline: %w", err)
	}

	sortedMaps := make(map[string][]string, len(keys))

	for i, key := range keys {
		members, err := sortCmds[i].Result()
		if err != nil {
			c.log.Error("Error fetching sorted members", zap.String("key", key), zap.Error(err))
			return nil, fmt.Errorf("error fetching sorted members for key %s: %w", key, err)
		}

		sortedMaps[key] = members
	}

	return sortedMaps, nil
}

// createFilterField creates the complete filter field by prefixing with the filter prefix.
func (c *RedisClient) createFilterField(suffix string) string {
	internalAttribute := c.getInternalField(suffix, nil)
	return fmt.Sprintf("%s:%s", c.prefixFilter, internalAttribute)
}

// createSortField creates the complete sorting field by prefixing with the sorting prefix.
func (c *RedisClient) createSortField(suffix string) string {
	internalAttribute := c.getInternalField(suffix, nil)
	return fmt.Sprintf("%s:%s", c.prefixSorting, internalAttribute)
}

// getInternalField maps the external field to an internal field based on provided mappings.
func (c *RedisClient) getInternalField(externalField string, customMappings map[string]string) string {
	fieldMappings := c.fieldMappings
	if customMappings != nil {
		fieldMappings = customMappings
	}

	internalField, ok := fieldMappings[externalField]
	if !ok {
		// If no mapping exists, use externalField directly
		return externalField
	}
	return internalField
}
