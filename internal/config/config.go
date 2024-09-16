package searchengine

import (
	"fmt"
	"os"
	"strings"

	envs "github.com/caarlos0/env/v11"
	"github.com/weissmedia/searchengine/internal/core"
)

// Config holds all configuration settings for the SearchEngine, including Redis settings,
// schema paths, and prefix values, with support for environment variables.
type Config struct {
	RedisHost        string              `env:"SEARCHENGINE_REDIS_HOST" envDefault:"localhost"`          // Hostname of the Redis server
	RedisPort        int                 `env:"SEARCHENGINE_REDIS_PORT" envDefault:"6379"`               // Port number of the Redis server
	RedisDB          int                 `env:"SEARCHENGINE_REDIS_DB" envDefault:"0"`                    // Redis database index to use
	RedisPassword    string              `env:"SEARCHENGINE_REDIS_PASSWORD" envDefault:""`               // Password for Redis authentication (if needed)
	UseSSL           bool                `env:"SEARCHENGINE_REDIS_USE_SSL" envDefault:"false"`           // Whether to use SSL/TLS for Redis connections
	DataPrefix       string              `env:"SEARCHENGINE_REDIS_DATA_PREFIX" envDefault:"data"`        // Prefix for data JSON storage
	FilterPrefix     string              `env:"SEARCHENGINE_REDIS_FILTER_PREFIX" envDefault:"filter"`    // Prefix for filter SET lists
	SortingPrefix    string              `env:"SEARCHENGINE_REDIS_SORTING_PREFIX" envDefault:"sorting"`  // Prefix for sorting ZSET lists
	SearchSchema     []core.SearchSchema `env:"-"`                                                       // Schema defining the fields and their types
	SearchIndexName  string              `env:"SEARCHENGINE_SEARCH_INDEX_NAME" envDefault:"idx"`         // RedisSearch index name
	SearchSchemaFile string              `env:"SEARCHENGINE_SEARCH_SCHEMA_FILE,required"`                // File path for the search schema
	NamespacePrefix  string              `env:"SEARCHENGINE_NAMESPACE_PREFIX" envDefault:"searchengine"` // Namespace prefix for organizing data
	LogLevel         string              `env:"SEARCHENGINE_LOG_LEVEL" envDefault:"info"`                // Log level for the logger
	ProfilerEnabled  bool                `env:"SEARCHENGINE_PROFILER_ENABLED" envDefault:"false"`
}

// NewConfig creates a new Config instance by loading environment variables and reading the search schema from a file.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	// Parse environment variables and store them in the Config struct
	if err := envs.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to parse environment variables: %w", err)
	}

	// Load the search schema from the specified file
	if err := loadSearchSchemaFromFile(cfg.SearchSchemaFile, &cfg.SearchSchema); err != nil {
		return nil, fmt.Errorf("failed to load search schema from file: %w", err)
	}

	return cfg, nil
}

// loadSearchSchemaFromFile reads the schema from a file and converts it into a slice of SearchSchema structs.
func loadSearchSchemaFromFile(path string, schema *[]core.SearchSchema) error {
	// Read file content
	fileData, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read search schema file: %w", err)
	}

	// Convert the JSON content into a SearchSchema slice
	toSchema, err := core.ConvertJSONStringToSchema(fileData)
	if err != nil {
		return fmt.Errorf("failed to convert JSON to schema: %w", err)
	}

	// Assign the schema data
	*schema = toSchema
	return nil
}

// Getter functions for the Config interface

// GetRedisAddr returns the Redis server address in the format "host:port".
func (c *Config) GetRedisAddr() string {
	return fmt.Sprintf("%s:%d", c.RedisHost, c.RedisPort)
}

// GetRedisDB returns the Redis database index to use.
func (c *Config) GetRedisDB() int {
	return c.RedisDB
}

// GetRedisHost returns the Redis server host.
func (c *Config) GetRedisHost() string {
	return c.RedisHost
}

// GetRedisPort returns the Redis server port.
func (c *Config) GetRedisPort() int {
	return c.RedisPort
}

// GetRedisPassword returns the Redis password for authentication.
func (c *Config) GetRedisPassword() string {
	return c.RedisPassword
}

// GetUseSSL returns whether to use SSL/TLS for Redis connections.
func (c *Config) GetUseSSL() bool {
	return c.UseSSL
}

// GetSearchIndexName returns the name of the search index used by RedisSearch.
func (c *Config) GetSearchIndexName() string {
	return c.SearchIndexName
}

// GetSearchSchema returns the search schema for the engine, detailing field definitions.
func (c *Config) GetSearchSchema() []core.SearchSchema {
	return c.SearchSchema
}

// GetNamespacePrefix returns the namespace prefix used for organizing Redis data.
func (c *Config) GetNamespacePrefix() string {
	return c.NamespacePrefix
}

// GetDataPrefix returns the full data prefix, combining the namespace and the data prefix.
func (c *Config) GetDataPrefix() string {
	return formatPrefix(c.NamespacePrefix, c.DataPrefix)
}

// GetFilterPrefix returns the full filter prefix, combining the namespace and the filter prefix.
func (c *Config) GetFilterPrefix() string {
	return formatPrefix(c.NamespacePrefix, c.FilterPrefix)
}

// GetSortingPrefix returns the full sorting prefix, combining the namespace and the sorting prefix.
func (c *Config) GetSortingPrefix() string {
	return formatPrefix(c.NamespacePrefix, c.SortingPrefix)
}

// GetProfilerEnabled indicates if the profiler is active.
func (c *Config) GetProfilerEnabled() bool {
	return c.ProfilerEnabled
}

// formatPrefix ensures that prefixes are combined without extra colons
func formatPrefix(namespace, prefix string) string {
	// Removes trailing colons from the namespace and leading colons from the prefix
	return strings.TrimSuffix(namespace, ":") + ":" + strings.TrimPrefix(prefix, ":")
}
