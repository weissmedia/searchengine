package core

import (
	"go.uber.org/zap"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

// AttributeType defines the type of attributes in the sorting map
type AttributeType int

const (
	IntType AttributeType = iota
	StringType
)

// SortResult holds the result of a sorting operation
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
	fields []SortField
	logger *zap.Logger // Store the logger here
}

// NewSortFieldList initializes a SortFieldList with a logger
func NewSortFieldList(logger *zap.Logger) *SortFieldList {
	return &SortFieldList{
		logger: logger,
	}
}

// AddSortField adds a new SortField to the SortFieldList
func (s *SortFieldList) AddSortField(name string, order string) {
	s.fields = append(s.fields, SortField{Name: name, Order: s.parseSortOrder(order)})
}

// SortFields returns a list of field names from the sorting fields
func (s *SortFieldList) SortFields() []string {
	names := make([]string, len(s.fields))
	for i, field := range s.fields {
		names[i] = field.Name
	}
	return names
}

// GetSortField retrieves a SortField by name from the SortFieldList
func (s *SortFieldList) GetSortField(name string) *SortField {
	for i := range s.fields {
		if s.fields[i].Name == name {
			return &s.fields[i]
		}
	}
	return nil
}

// Len returns the number of fields in the SortFieldList
func (s *SortFieldList) Len() int {
	return len(s.fields)
}

// parseSortOrder converts a string to a SortOrder type and logs if an unknown order is encountered
func (s *SortFieldList) parseSortOrder(order string) SortOrder {
	switch order {
	case "asc", "ASC":
		return Asc
	case "desc", "DESC":
		return Desc
	default:
		s.logger.Warn("Unknown sort order, setting default to ASC", zap.String("order", order))
		return Asc
	}
}

// BuildSortMap creates a sorting map from a list of strings, processed in parallel.
// Each string is expected to be in the format "value:key" and is split to create the map.
func (s *SortFieldList) BuildSortMap(list []string) (map[string]interface{}, AttributeType) {
	if len(list) == 0 {
		return nil, StringType // Return an empty map with a default type
	}

	// Determine the attribute type based on the first value in the list
	firstValue := strings.Split(list[0], ":")[0]
	attrType, parse := getAttributeType(firstValue)

	// Initialize the sorting map
	sortingMap := make(map[string]interface{}, len(list))

	var wg sync.WaitGroup
	chunkSize := (len(list) + runtime.NumCPU() - 1) / runtime.NumCPU()
	mu := &sync.Mutex{} // Mutex for safe access to the map in concurrent operations

	// Split the list into chunks and process each chunk in parallel
	for i := 0; i < len(list); i += chunkSize {
		end := i + chunkSize
		if end > len(list) {
			end = len(list)
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			localMap := make(map[string]interface{}, end-start)
			for _, s := range list[start:end] {
				parts := strings.Split(s, ":")
				if len(parts) != 2 {
					// Skip invalid entries and log a warning
					continue
				}
				localMap[parts[1]] = parse(parts[0])
			}

			// Lock the mutex before updating the global sorting map
			mu.Lock()
			for k, v := range localMap {
				sortingMap[k] = v
			}
			mu.Unlock()
		}(i, end)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	return sortingMap, attrType
}

// getAttributeType determines whether the attribute type is an integer or string
// and returns a parsing function for that type
func getAttributeType(val string) (AttributeType, func(string) interface{}) {
	if _, err := strconv.Atoi(val); err == nil {
		// If the value is an integer, return a parsing function for integers
		return IntType, func(s string) interface{} {
			v, _ := strconv.Atoi(s) // We ignore errors here since we know it's valid
			return v
		}
	}
	// Default to a string type if it's not an integer
	return StringType, func(s string) interface{} { return s }
}
