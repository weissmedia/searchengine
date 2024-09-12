package core

// SearchSchemaType is an enumeration for field types
type SearchSchemaType int

type SearchSchema struct {
	Path          string           `json:"-"`
	Name          string           `json:"name"`
	Type          SearchSchemaType `json:"type"`
	SearchOptions string           `json:"search_options"`
}

const (
	TextField SearchSchemaType = iota
	NumericField
)

const (
	TextSearchOptions    = "fuzzy, prefix, wildcard"
	NumericSearchOptions = "range"
)

// String method for FieldType to get the string representation
func (s SearchSchemaType) String() string {
	switch s {
	case TextField:
		return "text"
	case NumericField:
		return "numeric"
	default:
		return "unknown"
	}
}
