package backend

type SearchBackend interface {
	Search(query string) ([]string, error)
	DefineSchema(index string, schema []SchemaField) error
	DropSchema(index string) error
	RebuildSchema(index string, schema []SchemaField) error
}

// SchemaField beschreibt ein Feld in einem Suchschema
type SchemaField struct {
	Name string
	Type string // TEXT, NUMERIC, TAG, etc.
}
