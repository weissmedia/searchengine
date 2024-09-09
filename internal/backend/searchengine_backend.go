package backend

type SearchBackend interface {
	Search(query string) ([]string, error)
	DefineSchema() (bool, error)
	DropSchema() (bool, error)
	RebuildSchema() (bool, error)
}
