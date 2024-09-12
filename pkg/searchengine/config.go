package searchengine

// Config enthält die Konfigurationseinstellungen für die SearchEngine
type Config struct {
	RedisAddress   string
	UseRedisSearch bool
	SearchSchema   SearchSchema // Kennzeichnung, dass es sich um ein Suchschema handelt
	IndexName      string       // Anwender kann den Indexnamen setzen
}

func (c *Config) GetRedisAddr() string {
	return c.RedisAddress
}
func (c *Config) GetIndexName() string {
	return c.IndexName
}
func (c *Config) GetSearchSchema() SearchSchema {
	return c.SearchSchema
}
