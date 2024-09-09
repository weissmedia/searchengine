package redissearch

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

type Client struct {
	client *redis.Client
}

// NewClient erstellt eine neue Instanz des RedisSearchClient
func NewClient(redisAddr string) *Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	return &Client{
		client: rdb,
	}
}

// Search führt eine RedisSearch-Volltextsuche durch
func (r *Client) Search(index, query string) ([]string, error) {
	cmd := r.client.Do(context.Background(), "FT.SEARCH", index, query)
	res, err := cmd.Result()
	if err != nil {
		log.Printf("Error executing RedisSearch query: %v", err)
		return nil, err
	}

	// Konvertiere das Ergebnis in ein String-Array
	var results []string
	for _, item := range res.([]interface{}) {
		results = append(results, item.(string))
	}
	return results, nil
}

// FuzzySearch führt eine Fuzzy-Suche in RedisSearch durch
func (r *Client) FuzzySearch(index, query string) ([]string, error) {
	// Beispiel für Fuzzy-Logik: Anhängen von "~" an den Suchbegriff
	fuzzyQuery := query + "~"
	return r.Search(index, fuzzyQuery)
}
