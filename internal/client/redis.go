package client

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"sort"
	"strings"
)

type RedisClient struct {
	client        *redis.Client
	fieldMappings map[string]string
}

func NewRedisClient() *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	return &RedisClient{
		client: client,
	}
}

func (c *RedisClient) SMembersMap(ctx context.Context, key string) (map[string]struct{}, error) {
	res := c.client.SMembersMap(ctx, c.createFilterField(key))
	return res.Result()
}

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

		// Überprüfen, ob der Wert nicht ausgeschlossen werden soll
		if len(parts) >= patternIndex {
			attr := parts[patternIndex-1]

			// Nur ausschließen, wenn ein gültiger `excludeValue` gesetzt ist
			if excludeValue == "" || attr != excludeValue {
				fieldValues = append(fieldValues, attr)
			}
		} else {
			log.Printf("Pattern index %d out of range for key: %s", patternIndex, key)
		}
	}

	if err := iter.Err(); err != nil {
		log.Printf("Error iterating over filter fields: %v", err)
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

func (c *RedisClient) SUnionMap(ctx context.Context, keys ...string) (map[string]struct{}, error) {
	// Erstelle eine Kopie von keys
	keysCopy := make([]string, len(keys))
	copy(keysCopy, keys)

	// Aktualisiere die Schlüssel in der Kopie
	for i, key := range keysCopy {
		keysCopy[i] = c.createFilterField(key) // Filterfeld für jeden Schlüssel anwenden
	}

	// Führe den SUNION-Befehl mit der Kopie der Schlüssel aus
	result, err := c.client.SUnion(ctx, keysCopy...).Result()
	if err != nil {
		return nil, err
	}

	// Konvertiere das Ergebnis (eine Liste von Strings) in eine Map
	unionMap := make(map[string]struct{}, len(result))
	for _, item := range result {
		unionMap[item] = struct{}{}
	}

	return unionMap, nil
}

func (c *RedisClient) ZRangeIndexMap(ctx context.Context, keys []string) ([]map[string]int, error) {
	pipe := c.client.Pipeline()
	sortCmds := make([]*redis.StringSliceCmd, len(keys))
	for i, key := range keys {
		sortCmds[i] = pipe.ZRange(ctx, c.createSortField(key), 0, -1)
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		return nil, err
	}

	sortedIndexMaps := make([]map[string]int, len(keys))

	for i, key := range keys {
		members, err := sortCmds[i].Result()
		if err != nil {
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

func (c *RedisClient) ZRangeMap(ctx context.Context, keys []string) (map[string][]string, error) {
	pipe := c.client.Pipeline()
	sortCmds := make([]*redis.StringSliceCmd, len(keys))

	for i, key := range keys {
		sortCmds[i] = pipe.ZRange(ctx, c.createSortField(key), 0, -1)
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("error executing pipeline: %w", err)
	}

	sortedMaps := make(map[string][]string, len(keys))

	for i, key := range keys {
		members, err := sortCmds[i].Result()
		if err != nil {
			return nil, fmt.Errorf("error fetching sorted members for key %s: %w", key, err)
		}

		sortedMaps[key] = members
	}

	return sortedMaps, nil
}

func (c *RedisClient) createFilterField(suffix string) string {
	internalAttribute := c.getInternalField(suffix, nil)
	// todo: make it configurable as quickly as possible
	return fmt.Sprintf("opus.bdl.datapool:filter:%s", internalAttribute)
}

func (c *RedisClient) createSortField(suffix string) string {
	internalAttribute := c.getInternalField(suffix, nil)
	// todo: make it configurable as quickly as possible
	return fmt.Sprintf("opus.bdl.datapool:sorting:%s", internalAttribute)
}

func (c *RedisClient) getInternalField(externalField string, customMappings map[string]string) string {
	fieldMappings := c.fieldMappings
	if customMappings != nil {
		fieldMappings = customMappings
	}

	internalField, ok := fieldMappings[externalField]
	if !ok {
		// If no mapping exists, use clientField directly
		return externalField
	}
	return internalField
}
