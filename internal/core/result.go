package core

import (
	"bytes"
	"encoding/json"
	"github.com/weissmedia/searchengine/internal/profiler"
)

// ExecutionResult represents the result of a search execution, including the results, timings, and total execution time.
type ExecutionResult struct {
	ResultSet          []string                   `json:"ResultSet"`          // The list of results from the search query
	ResultCount        int                        `json:"ResultCount"`        // The number of results in ResultSet
	Timings            []profiler.OperationTiming `json:"Timings"`            // Timings for each operation
	TotalExecutionTime float64                    `json:"TotalExecutionTime"` // Total execution time of the search in milliseconds
	Log                []string                   `json:"Log"`
}

// ToJSON converts ExecutionResult to JSON and returns it as a string.
func (e *ExecutionResult) ToJSON() (string, error) {
	// Create a buffer to hold the output
	var buffer bytes.Buffer

	// Create a new JSON encoder and disable HTML escaping
	encoder := json.NewEncoder(&buffer)
	encoder.SetEscapeHTML(false)

	// Encode the ExecutionResult struct into JSON
	err := encoder.Encode(e)
	if err != nil {
		return "", err
	}

	// Return the JSON output as a string
	return buffer.String(), nil
}
