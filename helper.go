package pgmeta

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

func NamedGet(db *sqlx.DB, dest interface{}, query string, arg interface{}) (err error) {
	rows, err := db.NamedQuery(query, arg)
	if err != nil {
		return err
	}
	defer rows.Close()
	if rows.Next() {
		return rows.StructScan(dest)
	}
	return sql.ErrNoRows
}
