package backend

import (
	"github.com/weissmedia/searchengine/internal/core"
)

// Config is an interface that defines the necessary configuration methods
// for the backend to interact with the Redis instance and the search engine.
type Config interface {
	// GetRedisAddr returns the address of the Redis server in the "host:port" format.
	// This is used to establish a connection with the Redis server.
	GetRedisAddr() string

	// GetRedisDB returns the Redis database index to use.
	// This allows the user to select a specific Redis database (default is 0).
	GetRedisDB() int

	// GetRedisHost returns the host part of the Redis server address.
	// It extracts and returns the hostname or IP address from the full Redis address.
	GetRedisHost() string

	// GetRedisPort returns the port part of the Redis server address as an integer.
	// This extracts the port from the Redis address and converts it to an integer.
	GetRedisPort() int

	// GetSearchIndexName returns the name of the search index to be used by the search engine.
	// This index is used to store and query the RedisSearch data.
	GetSearchIndexName() string

	// GetSearchSchema returns the schema that defines the fields and their types for the search engine.
	// The schema outlines the structure of the searchable fields (e.g., text, numeric).
	GetSearchSchema() []core.SearchSchema

	// GetRedisPassword (Optional) returns the password for Redis authentication, if required.
	// This method provides the password for connecting to Redis instances with authentication enabled.
	GetRedisPassword() string

	// GetUseSSL (Optional) returns whether SSL/TLS should be used for Redis connections.
	// This is used to determine if the connection to Redis should be secured via SSL/TLS.
	GetUseSSL() bool

	// GetNamespacePrefix returns the namespace prefix used to organize Redis keys.
	// This is used to logically group keys by a common prefix, such as "searchengine".
	GetNamespacePrefix() string

	// GetDataPrefix returns the prefix used for storing JSON data in Redis.
	// This prefix is used to organize keys that store the main data entities.
	GetDataPrefix() string

	// GetFilterPrefix returns the prefix used for storing filter-related SET lists in Redis.
	// This prefix is used to organize filter lists that are utilized for query-based filtering operations.
	GetFilterPrefix() string

	// GetSortingPrefix returns the prefix used for storing sorted ZSET lists in Redis.
	// This prefix is used to manage sorted sets that help with sorting and ranking operations.
	GetSortingPrefix() string

	// GetProfilerEnabled returns whether the profiler is enabled.
	// When true, performance metrics are collected during query execution.
	GetProfilerEnabled() bool
}
