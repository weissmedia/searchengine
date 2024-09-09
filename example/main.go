package main

import (
	"fmt"
	"github.com/weissmedia/searchengine/internal/config"
	"github.com/weissmedia/searchengine/internal/query"
)

func main() {
	schema := []config.SearchSchema{
		{Name: "title", Type: "TEXT"},
		{Name: "description", Type: "TEXT"},
		{Name: "price", Type: "NUMERIC"},
		{Name: "category", Type: "TAG"},
	}

	// Benutzerdefinierte Config erstellen
	cfg := config.Config{
		RedisAddress: "localhost:6379", // pflicht
		IndexName:    "myIndex",        // optional
		SearchSchema: schema,
	}

	// Erstelle die SearchEngine und übergebe das Backend und die Config
	se := query.NewSearchEngine(cfg)

	// Recreate das Suchschema
	err := se.RecreateSchema()
	if err != nil {
		fmt.Println("Error recreating schema:", err)
		return
	}

	// Führe eine Suche durch
	result, err := se.Search("title:Redis")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Search result:", result)
}
