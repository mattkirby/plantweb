package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	db *sql.DB
}

func (s *Sqlite) Open(d string) error {
	log.Println("trying")
	db, err := sql.Open("sqlite3", d)
	if err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *Sqlite) Close() {
	s.db.Close()
}

func (s *Sqlite) Exec(ss string) error {
	_, err := s.db.Exec(ss)
	if err != nil {
		log.Printf("%q: %s\n", err, ss)
		return err
	}
	return nil
}

func (s *Sqlite) Begin(ss string, d []string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(ss)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for i, v := range d {
		_, err = stmt.Exec(i, fmt.Sprintf("%s%03d", v, i))
		if err != nil {
			return err
		}
	}

	////for i := 0; i < 100; i++ {
	////	_, err = stmt.Exec(i, fmt.Sprintf("こんにちは世界%03d", i))
	////	if err != nil {
	////		return err
	////	}
	////}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (s *Sqlite) Query(ss string) (result []string, err error) {
	rows, err := s.db.Query(ss)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			return
		}
		result = append(result, fmt.Sprintf("%d %s", id, name))
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}

func (s *Sqlite) Prepare(ss string) (result string, err error) {
	stmt, err := s.db.Prepare(ss)
	if err != nil {
		return
	}
	defer stmt.Close()
	var name string
	err = stmt.QueryRow("3").Scan(&name)
	if err != nil {
		return
	}
	result = name
	return
}
