package pgmeta

import "github.com/jmoiron/sqlx"

type Schema struct {
	Id    int64  `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Owner string `db:"owner" json:"owner"`
}

func ListSchemas(db *sqlx.DB) (rows []Schema, err error) {
	err = db.Select(&rows, QueryGetAllSchemas)
	return
}

func GetSchema(db *sqlx.DB, schemaName string) (schema Schema, err error) {
	err = db.Get(&schema, QueryGetSchema, schemaName)
	return
}
