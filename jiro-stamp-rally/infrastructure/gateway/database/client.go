package client

import "database/sql"

type DbClient interface {
	InsertSQL(tablename string, valueMap map[string]any) error
	SelectSQL(query string) (*sql.Rows, error)
}
