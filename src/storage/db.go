package storage

type Storage interface {
	Open() error
	Close()
	Exec(string) error
	Begin(string) error
	Query(string) error
	Prepare(string) error
}
