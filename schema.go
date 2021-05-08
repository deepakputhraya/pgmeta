package pgmeta

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"strings"
)

type Schema struct {
	Id    int64  `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Owner string `db:"owner" json:"owner"`
}

// TODO: Update Schema

func CreateSchema(db *sqlx.DB, schemaName string) (Schema, error) {
	if len(strings.TrimSpace(schemaName)) == 0 {
		return Schema{}, errors.New("schema name should not be blank")
	}
	_, err := db.Exec(fmt.Sprintf("CREATE SCHEMA %s", pq.QuoteIdentifier(schemaName)))
	if err != nil {
		return Schema{}, err
	}
	return GetSchema(db, schemaName)
}

func DeleteSchema(db *sqlx.DB, schemaName string) error {
	if len(strings.TrimSpace(schemaName)) == 0 {
		return errors.New("schema name should not be blank")
	}
	_, err := db.Exec(fmt.Sprintf("DROP SCHEMA %s", pq.QuoteIdentifier(schemaName)))
	return err
}

func ListSchemas(db *sqlx.DB) (rows []Schema, err error) {
	err = db.Select(&rows, QueryGetAllSchemas)
	return
}

func GetSchema(db *sqlx.DB, schemaName string) (schema Schema, err error) {
	err = db.Get(&schema, QueryGetSchema, schemaName)
	return
}
