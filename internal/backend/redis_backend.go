package backend

import (
	"github.com/weissmedia/searchengine/internal/redis"
	"github.com/weissmedia/searchengine/internal/redissearch"
)

// RedisBackend vereint Redis und RedisSearch in einem Backend
type RedisBackend struct {
	redisClient       *redis.Client
	redisSearchClient *redissearch.Client
	config            Config
}

// NewRedisBackend erstellt eine Instanz von RedisBackend
func NewRedisBackend(config Config) *RedisBackend {
	return &RedisBackend{
		redisClient:       redis.NewClient(config.GetRedisAddr()),
		redisSearchClient: redissearch.NewClient(config.GetRedisAddr()),
		config:            config,
	}
}

func (r *RedisBackend) DefineSchema() (bool, error) {
	return true, nil
}

// DropSchema löscht das RedisSearch-Schema
func (r *RedisBackend) DropSchema() (bool, error) {
	return true, nil
}

func (r *RedisBackend) RebuildSchema() (bool, error) {
	if ok, err := r.DropSchema(); ok {
		return false, err
	}
	return r.DefineSchema()
}

func (r *RedisBackend) Search(query string) ([]string, error) {
	//TODO implement me
	panic("implement me")
}
