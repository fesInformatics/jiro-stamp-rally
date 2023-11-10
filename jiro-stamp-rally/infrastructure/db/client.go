package client

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	database "github.com/fesInformatics/jiro-stamp-rally/infrastructure/gateway/database"
)

type DbClient struct {}

func (dc *DbClient) InsertSQL(tablename string, valueMap map[string]any) error {
	var (
		columns string = ""
		here string = ""
		values = make([]any, 0, len(valueMap))
	)
	db, err := sql.Open("mysql", "user:password@tcp(mysql:3306)/jiro-stamp-rally?charset=utf8mb4")
	if err != nil {
		return err
	}

	for k, v := range valueMap {
		if columns != "" {
			columns += ","
		} 
		columns += k
		if here == "" {
			here += "?"
		} else {
			here += ",?"
		}
		values = append(values, v)
	}

	ins, err := db.Prepare(fmt.Sprintf("INSERT INTO %s(%s) VALUES(%s)", tablename, columns, here))
	if err != nil {
		return err
	}
	fmt.Printf("バリュー:%v",values...)
	_, err = ins.Exec(values...)
	return err
}

func NewDbClient() database.DbClient {
	return &DbClient{}
}
