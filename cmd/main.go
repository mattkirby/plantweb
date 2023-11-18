package main

import (
	"github.com/mattkirby/plantweb/src/storage"
	"log"
)

func main() {
	//var db storage.Db = &storage.Postgresql{}
	var db storage.Db = &storage.Sqlite{}
	err := db.Open("./tmp/sqlite.db")
	//err := db.Open("user=mk host=localhost port=5432 dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table foo (id integer not null primary key, name text);
	delete from foo;
	`

	err = db.Exec(sqlStmt)
	if err != nil {
		log.Println(err)
	}

	ss := "insert into foo(id, name) values(?, ?)"
	var data []string
	data = append(data, "testthing", "testthing", "testthing", "testthing")
	err = db.Begin(ss, data)
	if err != nil {
		log.Println(err)
	}

	ss = "select id, name from foo"
	res, err := db.Query(ss)
	if err != nil {
		log.Println(err)
	}
	log.Println(res)

	ss = "select name from foo where id = ?"
	result, err := db.Prepare(ss)
	if err != nil {
		log.Println(err)
	}
	log.Println(result)

	ss = "delete from foo"
	err = db.Exec(ss)
	if err != nil {
		log.Println(err)
	}
}
