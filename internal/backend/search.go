package backend

import (
	"github.com/weissmedia/searchengine/internal/core"
	"golang.org/x/net/context"
)

type SearchBackend interface {
	// GetMap returns results for either single or multiple values on a field
	GetMap(ctx context.Context, field string, value interface{}) (map[string]struct{}, error)

	// GetMapExcluding retrieves all field values excluding the given one
	GetMapExcluding(ctx context.Context, field string, valueExclude interface{}) (map[string]struct{}, error)
	// GetFieldValuesExcluding retrieves field values excluding the given one
	GetFieldValuesExcluding(ctx context.Context, field string, valueExclude interface{}) ([]string, int, error)
	// SearchComparisonMap handles a comparison query and executes it using the appropriate operator
	SearchComparisonMap(field string, operator core.ComparisonOperator, value interface{}) (map[string]struct{}, error)

	// SearchRangeMap constructs and executes a range query in Redis
	SearchRangeMap(field string, min, max interface{}) (map[string]struct{}, error)
	// SearchFuzzyMap constructs and executes a fuzzy search query in Redis
	SearchFuzzyMap(field, value string) (map[string]struct{}, error)
	SearchWildcardMap(field, value string) (map[string]struct{}, error)

	// GetSortedFieldValuesMap retrieves sorted field values using Redis zRange
	GetSortedFieldValuesMap(ctx context.Context, fields []string) (<-chan core.SortResult, error)
}
