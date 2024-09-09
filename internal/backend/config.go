package backend

import "github.com/weissmedia/searchengine/internal/schema"

type Config interface {
	GetRedisAddr() string
	GetIndexName() string
	GetSearchSchema() schema.SearchSchema
}

// SearchSchema definiert das Schema für Suchabfragen
type SearchSchema struct {
	Fields []SchemaField
}

// SchemaField beschreibt ein Feld in einem Suchschema
type SchemaField struct {
	Name string
	Type string // TEXT, NUMERIC, TAG, etc.
}
