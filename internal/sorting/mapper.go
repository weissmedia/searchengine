package sorting

import (
	"github.com/weissmedia/searchengine/internal/core"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

// Mapper defines the interface for creating sorting maps
type Mapper interface {
	CreateSortingMap(list []string) (map[string]interface{}, core.AttributeType)
}

// DefaultMapper is the default implementation of SortingMapper
type DefaultMapper struct{}

func NewMapper() *DefaultMapper {
	return &DefaultMapper{}
}

// CreateSortingMap creates a sorting map from a list of strings and a parsing function in parallel
func (s *DefaultMapper) CreateSortingMap(list []string) (map[string]interface{}, core.AttributeType) {
	if len(list) == 0 {
		return nil, core.StringType // Or an appropriate default value
	}

	firstValue := strings.Split(list[0], ":")[0]
	attrType, parse := getAttributeType(firstValue)
	sortingMap := make(map[string]interface{}, len(list))

	var wg sync.WaitGroup
	chunkSize := (len(list) + runtime.NumCPU() - 1) / runtime.NumCPU()

	mu := &sync.Mutex{} // Mutex for safe access to the map

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
				localMap[parts[1]] = parse(parts[0])
			}
			mu.Lock()
			for k, v := range localMap {
				sortingMap[k] = v
			}
			mu.Unlock()
		}(i, end)
	}

	wg.Wait()
	return sortingMap, attrType
}

// getAttributeType function
func getAttributeType(val string) (core.AttributeType, func(string) interface{}) {
	if _, err := strconv.Atoi(val); err == nil {
		return core.IntType, func(s string) interface{} { v, _ := strconv.Atoi(s); return v }
	}
	return core.StringType, func(s string) interface{} { return s }
}
