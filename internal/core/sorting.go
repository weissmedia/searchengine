package core

import "log"

type AttributeType int

const (
	IntType AttributeType = iota
	StringType
)

type SortResult struct {
	Index        int
	Field        string
	OrderMap     map[string]interface{}
	OrderMapType AttributeType
	Err          error
}

// SortOrder defines the sorting order
type SortOrder int

const (
	Asc SortOrder = iota
	Desc
)

// SortField represents a single attribute to be sorted by
type SortField struct {
	Name  string
	Order SortOrder
}

// SortFieldList contains a list of sorting fields
type SortFieldList struct {
	Fields []SortField
}

// AddSortField adds a new SortField to the SortFieldList
func (s *SortFieldList) AddSortField(name string, order string) {
	s.Fields = append(s.Fields, SortField{Name: name, Order: parseSortOrder(order)})
}

// SortFields returns a list of field names from the sorting fields
func (s *SortFieldList) SortFields() []string {
	names := make([]string, len(s.Fields))
	for i, field := range s.Fields {
		names[i] = field.Name
	}
	return names
}

func (s *SortFieldList) GetSortField(name string) *SortField {
	for i := range s.Fields {
		if s.Fields[i].Name == name {
			return &s.Fields[i]
		}
	}
	return nil
}

func (s *SortFieldList) Len() int {
	return len(s.Fields)
}

// ParseSortOrder converts a string to a SortOrder type
func parseSortOrder(order string) SortOrder {
	switch order {
	case "asc", "ASC":
		return Asc
	case "desc", "DESC":
		return Desc
	default:
		log.Printf("unknown sort order: %s. set default order asc.", order)
		return Asc
	}
}
