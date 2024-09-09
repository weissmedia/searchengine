package backend

import (
	"github.com/weissmedia/searchengine/internal/redis"
	"github.com/weissmedia/searchengine/internal/redissearch"
)

// RedisBackend vereint Redis und RedisSearch in einem Backend
type RedisBackend struct {
	redisClient       *redis.Client
	redisSearchClient *redissearch.Client
}

// NewRedisBackend erstellt eine Instanz von RedisBackend
func NewRedisBackend(redisAddr string) *RedisBackend {
	return &RedisBackend{
		redisClient:       redis.NewClient(redisAddr),
		redisSearchClient: redissearch.NewClient(redisAddr),
	}
}

func (r *RedisBackend) DefineSchema(index string, schema []SchemaField) error {
	return nil
}

// DropSchema löscht das RedisSearch-Schema
func (r *RedisBackend) DropSchema(index string) error {
	return nil
}

// RecreateSchema löscht das Schema und erstellt es neu
func (r *RedisBackend) RecreateSchema(index string, schema []SchemaField) error {
	err := r.DropSchema(index)
	if err != nil {
		return err
	}
	return r.DefineSchema(index, schema)
}

func (r *RedisBackend) RebuildSchema(index string, schema []SchemaField) error {
	//TODO implement me
	panic("implement me")
}

func (r *RedisBackend) Search(query string) ([]string, error) {
	//TODO implement me
	panic("implement me")
}
