package pgmeta

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type PrimaryKey struct {
	Schema    string `db:"schema" json:"schema"`
	TableName string `db:"table_name" json:"table_name"`
	TableId   int64  `db:"table_id" json:"table_id"`
	Name      string `db:"name" json:"name"`
}

func ListPrimaryKeys(db *sqlx.DB, schema string, tableName string) ([]PrimaryKey, error) {
	rawRows, err := db.NamedQuery(QueryPrimaryKeys, map[string]interface{}{
		"tableName": tableName,
		"schema":    schema,
	})
	if err != nil {
		return nil, err
	}
	var rows []PrimaryKey
	for rawRows.Next() {
		var row PrimaryKey
		err := rawRows.StructScan(&row)
		if err != nil {
			return nil, err
		}
		rows = append(rows, row)
	}
	return rows, nil
}

type Column struct {
	ColumnName      string         `db:"column_name" json:"column_name"`
	DataType        string         `db:"data_type" json:"data_type"`
	Schema          string         `db:"table_schema" json:"schema"`
	TableName       string         `db:"table_name" json:"table_name"`
	OrdinalPosition string         `db:"ordinal_position" json:"ordinal_position"`
	ColumnDefault   sql.NullString `db:"column_default" json:"default_value"`
	IsUpdatable     bool           `db:"is_updatable" json:"is_updatable"`
	IsNullable      bool           `db:"is_nullable" json:"is_nullable"`
	Id              int64          `db:"dtd_identifier" json:"id"`
}

func ListColumns(db *sqlx.DB, schema string, tableName string) ([]Column, error) {
	rawRows, err := db.NamedQuery(QueryTableColumns, map[string]interface{}{
		"tableName": tableName,
		"schema":    schema,
	})
	if err != nil {
		return nil, err
	}
	var rows []Column
	for rawRows.Next() {
		var row Column
		err := rawRows.StructScan(&row)
		if err != nil {
			return nil, err
		}
		rows = append(rows, row)
	}
	return rows, nil
}

type Table struct {
	Id                        string         `db:"id" json:"id"`
	Schema                    string         `db:"schema" json:"schema"`
	Name                      string         `db:"name" json:"name"`
	HasRowLevelSecurity       bool           `db:"rls_security" json:"rls_security"`
	IsRowLevelSecurityEnabled bool           `db:"rls_enabled" json:"rls_enabled"`
	Bytes                     int64          `db:"bytes" json:"bytes"`
	Size                      string         `db:"size" json:"size"`
	Comment                   sql.NullString `db:"comment" json:"comment"`
	EstimatedLiveRows         int64          `db:"live_rows_estimate" json:"estimated_live_rows"`
	EstimatedDeadRows         int64          `db:"dead_rows_estimate" json:"estimated_dead_rows"`
}

func ListTables(db *sqlx.DB, schema string) ([]Table, error) {
	rawRows, err := db.NamedQuery(QueryListTables, map[string]interface{}{
		"schema": schema,
	})
	if err != nil {
		return nil, err
	}
	var rows []Table
	for rawRows.Next() {
		var row Table
		err := rawRows.StructScan(&row)
		if err != nil {
			return nil, err
		}
		rows = append(rows, row)
	}
	return rows, nil
}

func GetTable(db *sqlx.DB, schema string, tableName string) (table Table, err error) {
	err = NamedGet(db, &table, QueryGetTable, map[string]interface{}{
		"schema":    schema,
		"tableName": tableName,
	})
	return
}
