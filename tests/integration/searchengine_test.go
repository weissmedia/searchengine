package integration_test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/weissmedia/searchengine/internal/log"
	"github.com/weissmedia/searchengine/internal/search"
	"github.com/weissmedia/searchengine/pkg/searchengine"
	"go.uber.org/zap"
	"io"
	"os"
	"reflect"
	"testing"
)

var ctx = context.Background()
var logger *zap.Logger
var engine *search.Engine

func init() {
	logger = log.GetLogger()

	// Load the configuration
	cfg, err := searchengine.NewConfig()
	if err != nil {
		logger.Error("Error loading config", zap.Error(err))
		return
	}
	// Initialize the search engine with the configuration and logger
	engine = searchengine.NewEngine(cfg, logger)

	// Update search index and handle any errors
	isUpdated, err := engine.Backend.UpdateSearchIndex(cfg.GetSearchIndexName())
	if err != nil {
		logger.Error("Error updating search index", zap.Error(err))
		return
	}

	// Log if the search index was successfully updated
	if isUpdated {
		logger.Info("Search index was updated successfully", zap.String("index", cfg.GetSearchIndexName()))
	} else {
		logger.Warn("Search index was not updated", zap.String("index", cfg.GetSearchIndexName()))
	}
}

// Load JSON from a file into a map or slice
func loadJSON(filePath string, target interface{}) error {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	return json.Unmarshal(byteValue, target)
}

func TestMain(m *testing.M) {
	// Verbinde dich zu Redis
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	// Lade redisData aus der JSON-Datei
	redisData := make(map[string][]string)
	err := loadJSON("./data/filter.json", &redisData)
	if err != nil {
		logger.Fatal("Failed to load filter data:", zap.Error(err))
	}

	// Lade redisJSON aus der JSON-Datei
	redisJSON := make([]map[string]interface{}, 0)
	err = loadJSON("./data/data.json", &redisJSON)
	if err != nil {
		logger.Fatal("Failed to load json data:", zap.Error(err))
	}

	// Lade sortingData aus der JSON-Datei
	sortingData := make(map[string][]string)
	err = loadJSON("./data/sorting.json", &sortingData)
	if err != nil {
		logger.Fatal("Failed to load sorting data:", zap.Error(err))
	}

	// Import data into Redis
	for key, values := range redisData {
		err := rdb.SAdd(ctx, "opus.bdl.datapool:filter:"+key, values).Err()
		if err != nil {
			logger.Fatal("Failed to add key", zap.String("key", key), zap.Error(err))
		}
		logger.Info("Key added with values", zap.String("key", key), zap.Any("values", values))
	}

	// Insert JSON data into Redis
	for _, doc := range redisJSON {
		// Access the "id" field from the nested "data" field to construct the key
		dataField := doc["data"].(map[string]interface{})
		id := dataField["id"].(string)

		// Create a Redis key using the "id" from the "data" field
		key := fmt.Sprintf("opus.bdl.datapool:data:%s", id)

		// Serialize the entire document (including the "data" field)
		jsonData, err := json.Marshal(doc)
		if err != nil {
			logger.Fatal("Error serializing JSON", zap.Error(err))
		}

		// Store the complete document in Redis, including the "data" field
		err = rdb.Do(ctx, "JSON.SET", key, "$", string(jsonData)).Err()
		if err != nil {
			logger.Fatal("Error setting JSON in Redis", zap.String("key", key), zap.Error(err))
		}

		// Log the stored JSON data
		logger.Info("Complete JSON data with 'data' field stored in Redis", zap.String("key", key), zap.ByteString("data", jsonData))
	}

	// Store sorting data in Redis as sorted sets with ZADD
	for key, values := range sortingData {
		for _, member := range values {
			err := rdb.ZAdd(ctx, "opus.bdl.datapool:"+key, redis.Z{
				Score:  0,      // The score represents the position in the list
				Member: member, // The list element
			}).Err()
			if err != nil {
				logger.Fatal("Failed to add member", zap.String("member", member), zap.String("key", key), zap.Error(err))
			}
			logger.Info("Member added to sorted set", zap.String("member", member), zap.String("key", key))
		}
	}
	logger.Info("Data successfully imported.")

	// Starte die Tests
	code := m.Run()

	// Sync den Logger, um gepufferte Log-Einträge zu schreiben
	log.SyncLogger()

	// Beende die Tests mit dem Rückgabewert der Tests
	os.Exit(code)
}

func TestQueries(t *testing.T) {
	logger.Info("Starting test for search queries")
	// Definiere Suchanfragen mit den erwarteten Ergebnissen
	tests := []struct {
		query       string
		expected    []string
		checkOrder  bool
		expectError bool
	}{
		// Grundlegende Tests
		{"keyA = 'val1'", []string{"x", "y", "z"}, false, false},
		{"keyA = 'val2' OR keyB = 'val3'", []string{"r", "q", "b", "p", "a", "c"}, false, false},
		{"keyA = 'val2' AND keyC = 'val2'", []string{"c"}, false, false},
		{"keyC IN ('val1', 'val2')", []string{"p", "q", "r", "s", "t", "u", "c", "x"}, false, false},

		// Grundlegende Tests für !=
		{"keyA != 'val1'", []string{"a", "b", "c", "d", "e", "f"}, false, false},
		{"keyB != 'val1'", []string{"k", "l", "p", "q", "j", "c", "r"}, false, false},
		{"keyC != 'val1'", []string{"s", "t", "u", "v", "w", "x", "c"}, false, false},
		{"keyD != 'val1'", []string{"f", "h", "e", "x", "i", "u", "g"}, false, false},

		// Kombinationen mit AND und OR
		{"keyA != 'val1' OR keyB = 'val2'", []string{"f", "k", "l", "c", "e", "d", "b", "a", "j"}, false, false},
		{"keyA != 'val1' OR keyC = 'val3'", []string{"a", "b", "c", "d", "e", "f", "v", "w", "x"}, false, false},
		{"keyA != 'val1' OR keyC != 'val3'", []string{"r", "f", "u", "t", "p", "q", "c", "b", "e", "s", "a", "d", "x"}, false, false},
		{"keyA != 'val1' AND (keyB = 'val2' OR keyC = 'val2')", []string{"c"}, false, false},

		// Verschachtelte Bedingungen
		{"(keyA != 'val1' AND keyC = 'val3') OR keyB = 'val3'", []string{"r", "q", "p"}, false, false},
		{"keyA != 'val1' AND keyB != 'val2' AND keyC = 'val3'", []string{}, false, false},

		// Tests mit Sortierung
		{"keyA = 'val1' SORT BY keyA DESC", []string{"y", "z", "x"}, true, false},
		{"keyA = 'val1' SORT BY keyA ASC", []string{"x", "z", "y"}, true, false},
		{"keyB = 'val3' SORT BY keyB ASC, keyC DESC", []string{"q", "r", "p"}, true, false},
		{"keyC = 'val1' SORT BY keyC ASC, keyB DESC", []string{"q", "r", "p"}, true, false},
		{"keyC = 'val3' AND keyD = 'val2' SORT BY keyC ASC, keyD DESC", []string{"x"}, true, false},

		// Tests mit Limit und Offset
		{"keyA = 'val1' SORT BY keyA ASC LIMIT 2", []string{"x", "z"}, true, false},
		{"keyA = 'val1' SORT BY keyA DESC LIMIT 2", []string{"y", "z"}, true, false},
		{"keyA = 'val1' SORT BY keyA DESC LIMIT 1 OFFSET 1", []string{"z"}, true, false},
		{"keyA = 'val1' SORT BY keyA DESC LIMIT 2 OFFSET 1", []string{"z", "x"}, true, false},
		{"keyC IN ('val1', 'val2') SORT BY keyC ASC LIMIT 2", []string{"x", "c"}, true, false},

		// Tests mit verschachtelten Klammerungen
		{"(keyA = 'val1' OR keyB = 'val3') AND keyC='val2'", []string{"x"}, false, false},
		{"((keyA = 'val2' AND keyC='val2') OR (keyB = 'val3' AND data_num >= 10)) AND data_keyE=[0 100]", []string{"c"}, false, false},
		{"((keyA = 'val1' OR keyB = 'val3') AND (keyC = 'val2' OR keyD = 'val2')) AND data_keyE=[0 100] AND keyC='val2'", []string{"x"}, false, false},

		// Verschachtelte Klammerungen mit Sortierung
		{"((keyA = 'val1' OR keyB = 'val3') AND keyC='val1') SORT BY keyC ASC", []string{"q", "r", "p"}, true, false},
		{"(keyA = 'val1' AND (keyC='val1' OR keyD='val2')) SORT BY keyC DESC", []string{"x"}, true, false},
		{"(keyA = 'val2' AND (keyC='val1' OR keyD='val3')) SORT BY keyD ASC, keyC DESC", []string{}, true, false},
		{"(keyA = 'val1' OR (keyB = 'val3' AND data_num >= 10)) SORT BY keyC DESC", []string{"y", "z", "x"}, true, false},
		{"(keyA = 'val1' OR (keyB = 'val3' AND data_num >= 10)) SORT BY keyF DESC LIMIT 2 OFFSET 1", []string{"z", "x"}, true, false},

		// Kombination von Limit, Offset und verschachtelten Klammerungen
		{"((keyA = 'val1' OR keyB = 'val3') AND keyC='val1') SORT BY keyC DESC LIMIT 1 OFFSET 1", []string{"r"}, true, false},
		{"(keyA = 'val1' AND (keyC='val1' OR keyD='val2')) SORT BY keyC ASC LIMIT 2 OFFSET 1", []string{}, true, false},
		{"(((keyA = 'val1') OR (keyB = 'val3')) AND (keyC='val1' OR keyD='val2')) SORT BY keyC ASC, keyD DESC LIMIT 1", []string{"x"}, true, false},

		// Fuzzy Matching
		{"data_keyC ~ 'val1'", []string{"p", "q", "r"}, false, false},
		{"data_keyD ~ 'val3'", []string{"g", "h", "i"}, false, false},

		// Wildcard Matching
		{"data_keyA = '*al1'", []string{"x", "y", "z"}, false, false},
		{"data_keyF = 'va*'", []string{"j", "k", "l", "o", "q", "p", "r"}, false, false},
		{"data_keyD = '*val3*'", []string{"g", "h", "i"}, false, false},

		// IN-Operator mit Wildcards und Fuzzy-Matching
		//{"keyA IN ('val1', 'val2', '~val3')", []string{"x", "y", "z", "a", "b", "c", "d", "e", "f"}, false, true},
		{"keyD IN ('va*', 'val2', 'val3')", []string{"d", "e", "f", "g", "h", "i"}, false, true},

		// Verschachtelte Abfragen mit Fuzzy und Wildcard
		{"((keyA = 'val1' OR keyB = 'val3') AND data_keyC ~ 'val1')", []string{"p", "q", "r"}, false, false},
		{"(keyD = 'val2' AND data_keyC = '*al1')", []string{"u"}, false, false},
		{"(keyA = 'val1' AND data_keyC = '*al1' OR data_keyD ~ 'val3')", []string{"i", "g", "h"}, false, false},

		// Sortierung mit Wildcards und Fuzzy-Matching
		{"data_keyC = '*al1' SORT BY keyC DESC", []string{"s", "t", "u"}, true, false},
		{"data_keyD ~ 'val3' SORT BY keyD ASC", []string{"i", "h", "g"}, true, false},
		{"keyD IN ('val2', '*val3') SORT BY keyD ASC, keyC DESC", []string{"f", "g", "h"}, true, true},

		// Weitere Tests
		{"keyA = 'val2' AND (keyB = 'val2' OR keyC = 'val3')", []string{"c"}, false, false},
		{"keyA != 'val2' AND (keyB != 'val1' OR keyC != 'val1')", []string{"x"}, false, false},
		{"keyA = 'val1' AND data_num >= 10", []string{"x"}, false, false},
		{"keyB = 'val3' AND data_num <= 5", []string{}, false, false},
		{"data_keyD ~ 'val'", []string{"g", "h", "i"}, false, false},
	}

	// Führe die Tests aus
	for _, test := range tests {
		runTest(t, test.query, test.expected, test.checkOrder, test.expectError)
	}
}

// Funktion zum Ausführen einer Testabfrage und Vergleichen des Ergebnisses mit den erwarteten Werten
func runTest(t *testing.T, queryStr string, expected []string, checkOrder bool, expectError bool) {
	logger.Info("Query:", zap.String("query", queryStr))
	defer func() {
		if r := recover(); r != nil {
			if expectError {
				t.Logf("Test passed for query: %s (expected error occurred)", queryStr)
			} else {
				t.Errorf("Test failed for query: %s\nUnexpected error: %v", queryStr, r)
			}
		}
	}()
	// Execute a search query
	resultList, err := engine.Search(ctx, queryStr)

	if err != nil {
		if expectError {
			t.Logf("Test passed for query: %s (expected execution error occurred)", queryStr)
		} else {
			t.Fatalf("Execution error: %v", err)
		}
		return
	}

	// Falls kein Fehler erwartet wurde, aber einer aufgetreten ist
	if expectError {
		t.Errorf("Test failed for query: %s\nExpected an error, but query executed successfully", queryStr)
		return
	}

	t.Logf("Result from Visitor: %v", resultList.ResultSet)

	// Falls die Reihenfolge wichtig ist, vergleiche direkt
	if checkOrder {
		if !reflect.DeepEqual(resultList.ResultSet, expected) {
			t.Errorf("Test failed for query: %s\nExpected: %v\nGot: %v", queryStr, expected, resultList)
		} else {
			t.Logf("Test passed for query: %s", queryStr)
		}
	} else {
		// Vergleiche das Ergebnis als Set (ungeordneter Vergleich)
		if !equalAsSets(resultList.ResultSet, expected) {
			t.Errorf("Test failed for query: %s\nExpected: %v\nGot: %v", queryStr, expected, resultList)
		} else {
			t.Logf("Test passed for query: %s", queryStr)
		}
	}
}

// equalAsSets vergleicht zwei Slices, ohne die Reihenfolge zu berücksichtigen
func equalAsSets(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	// Map für die Elemente von 'a'
	set := make(map[string]struct{}, len(a))
	for _, v := range a {
		set[v] = struct{}{}
	}

	// Überprüfe, ob alle Elemente von 'b' in 'a' enthalten sind
	for _, v := range b {
		if _, ok := set[v]; !ok {
			return false
		}
	}

	return true
}
