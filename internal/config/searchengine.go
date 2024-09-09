package config

// SearchSchema definiert das Schema für Suchabfragen
type SearchSchema struct {
	Fields []SchemaField
}

// SchemaField beschreibt ein einzelnes Feld im Schema
type SchemaField struct {
	Name string
	Type string // TEXT, NUMERIC, TAG, etc.
}

// Config enthält die Konfigurationseinstellungen für die SearchEngine
type Config struct {
	RedisAddress   string
	UseRedisSearch bool
	SearchSchema   SearchSchema // Kennzeichnung, dass es sich um ein Suchschema handelt
	IndexName      string       // Anwender kann den Indexnamen setzen
}
