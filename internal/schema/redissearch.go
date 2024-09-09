package schema

// SearchSchema definiert das Schema für Suchabfragen
type SearchSchema struct {
	Fields []Field
}

// Field beschreibt ein einzelnes Feld im Schema
type Field struct {
	Name string
	Type string // TEXT, NUMERIC, TAG, etc.
}
