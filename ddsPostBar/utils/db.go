package utils

import (
	"database/sql"
	_"go-sql-driver/mysql"

)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "root:Dadan1223@tcp(localhost:3306)/dd_s_postbar")
	if err != nil {
		panic(err.Error())
	}
}

