package storage

type Db interface {
	Open(string) error
	Close()
	Exec(string) error
	Begin(string, []string) error
	Query(string) ([]string, error)
	Prepare(string) (string, error)
}
