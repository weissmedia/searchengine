package core

import (
	"encoding/json"
	"fmt"
	"strings"
)

// SearchSchemaType defines the type of the field in the search schema
type SearchSchemaType int

// Constants for SearchSchemaType
const (
	TextField SearchSchemaType = iota
	NumericField
)

// SearchSchema holds the generated schema for each field
type SearchSchema struct {
	Path          string           `json:"-"`
	Name          string           `json:"name"`
	Type          SearchSchemaType `json:"type"`
	SearchOptions string           `json:"search_options"`
}

// InputSchema defines the structure for the input schema
type InputSchema map[string]map[string]string

// ConvertInputToSchema converts an InputSchema struct into a list of SearchSchema structs
func ConvertInputToSchema(input InputSchema) ([]SearchSchema, error) {
	var schema []SearchSchema
	for index, fields := range input {
		for fieldName, fieldType := range fields {
			schemaType, err := mapFieldType(fieldType)
			if err != nil {
				return nil, fmt.Errorf("error processing field '%s' in '%s': %v", fieldName, index, err)
			}
			schema = append(schema, SearchSchema{
				Path:          fmt.Sprintf("$.%s.%s", index, fieldName),
				Name:          fmt.Sprintf("%s_%s", index, fieldName), // Prefix with the index for RedisSearch
				Type:          schemaType,
				SearchOptions: determineSearchOptions(schemaType),
			})
		}
	}
	return schema, nil
}

// ConvertJSONStringToSchema converts JSON data (string or byte slice) into a list of SearchSchema structs.
func ConvertJSONStringToSchema(jsonData interface{}) ([]SearchSchema, error) {
	var input InputSchema
	var jsonBytes []byte
	var err error

	// Handle different types of input: string or []byte
	switch data := jsonData.(type) {
	case string:
		jsonBytes = []byte(data) // Convert string to byte slice
	case []byte:
		jsonBytes = data
	default:
		return nil, fmt.Errorf("unsupported json data type: %T", jsonData)
	}

	// Unmarshal the JSON data into the InputSchema structure
	if err = json.Unmarshal(jsonBytes, &input); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	// Call ConvertInputToSchema to transform it into []SearchSchema
	return ConvertInputToSchema(input)
}

// mapFieldType maps the string field type to a SearchSchemaType
func mapFieldType(fieldType string) (SearchSchemaType, error) {
	switch strings.ToLower(fieldType) {
	case "text":
		return TextField, nil
	case "numeric":
		return NumericField, nil
	default:
		return -1, fmt.Errorf("unsupported field type: %s", fieldType)
	}
}

// determineSearchOptions provides search options based on the field type
func determineSearchOptions(fieldType SearchSchemaType) string {
	switch fieldType {
	case TextField:
		return "fuzzy, prefix, wildcard"
	case NumericField:
		return "range"
	default:
		return "unknown"
	}
}

// MarshalToJSON converts a SearchSchema slice to JSON
func MarshalToJSON(schema []SearchSchema) (string, error) {
	jsonData, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error marshalling schema to JSON: %v", err)
	}
	return string(jsonData), nil
}
