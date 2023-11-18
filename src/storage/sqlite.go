package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Db *sql.DB
}

func (s *Sqlite) Open() error {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		return err
	}
	s.Db = db
	return nil
}

func (s *Sqlite) Close() {
	s.Db.Close()
}

func (s *Sqlite) Exec(ss string) error {
	_, err := s.Db.Exec(ss)
	if err != nil {
		log.Printf("%q: %s\n", err, ss)
		return err
	}
	return nil
}

func (s *Sqlite) Begin(ss string) error {
	tx, err := s.Db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(ss)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for i := 0; i < 100; i++ {
		_, err = stmt.Exec(i, fmt.Sprintf("こんにちは世界%03d", i))
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (s *Sqlite) Query(ss string) error {
	rows, err := s.Db.Query(ss)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			return err
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}

func (s *Sqlite) Prepare(ss string) error {
	stmt, err := s.Db.Prepare(ss)
	if err != nil {
		return err
	}
	defer stmt.Close()
	var name string
	err = stmt.QueryRow("3").Scan(&name)
	if err != nil {
		return err
	}
	fmt.Println(name)
	return nil
}
