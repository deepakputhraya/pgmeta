package main

import (
	"fmt"
	"github.com/deepakputhraya/pgmeta"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	db, err := sqlx.Connect("postgres", "postgres://postgres:postgres@localhost:5432/waitlist?sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	db = db.Unsafe()

	tables, err := pgmeta.ListTables(db, "public")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tables - ")
	for _, data := range tables {
		fmt.Println(data)
	}

	columns, err := pgmeta.ListColumns(db, "public", "user")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Columns - ")
	for _, data := range columns {
		fmt.Println(data)
	}

	primaryKeys, err := pgmeta.ListPrimaryKeys(db, "public", "user")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Primary keys - ")
	for _, data := range primaryKeys {
		fmt.Println(data)
	}
}
