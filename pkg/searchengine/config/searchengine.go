package config

import (
	"github.com/weissmedia/searchengine/internal/schema"
)

// Config enthält die Konfigurationseinstellungen für die SearchEngine
type Config struct {
	RedisAddress   string
	UseRedisSearch bool
	SearchSchema   schema.SearchSchema // Kennzeichnung, dass es sich um ein Suchschema handelt
	IndexName      string              // Anwender kann den Indexnamen setzen
}

func (c *Config) GetRedisAddr() string {
	return c.RedisAddress
}
func (c *Config) GetIndexName() string {
	return c.IndexName
}
func (c *Config) GetSearchSchema() schema.SearchSchema {
	return c.SearchSchema
}
