package backend

import (
	"github.com/weissmedia/searchengine/pkg/searchengine"
)

type Config interface {
	GetRedisAddr() string
	GetIndexName() string
	GetSearchSchema() searchengine.SearchSchema
}
