package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

type Client struct {
	client *redis.Client
}

// NewClient erstellt eine neue Instanz von RedisClient mit der angegebenen Adresse
func NewClient(redisAddr string) *Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	return &Client{
		client: rdb,
	}
}

// Get holt einen Wert aus Redis basierend auf dem Schlüssel
func (r *Client) Get(key string) ([]string, error) {
	res, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		log.Printf("Failed to get key %s: %v", key, err)
		return nil, err
	}
	return []string{res}, nil
}

// Keys holt alle passenden Schlüssel basierend auf einem Muster
func (r *Client) Keys(pattern string) ([]string, error) {
	res, err := r.client.Keys(context.Background(), pattern).Result()
	if err != nil {
		log.Printf("Failed to get keys for pattern %s: %v", pattern, err)
		return nil, err
	}
	return res, nil
}

// Search führt eine Volltextsuche durch (falls RedisSearch genutzt wird)
func (r *Client) Search(index, query string) ([]string, error) {
	// Beispielcode für RedisSearch
	cmd := r.client.Do(context.Background(), "FT.SEARCH", index, query)
	res, err := cmd.Result()
	if err != nil {
		log.Printf("Failed to execute search on index %s with query %s: %v", index, query, err)
		return nil, err
	}

	// Konvertiere die Ergebnisse in ein String-Array
	var results []string
	for _, item := range res.([]interface{}) {
		results = append(results, item.(string))
	}

	return results, nil
}
