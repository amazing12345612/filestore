package mysql

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
)

func TestMysql(m *testing.T) {
	db, _ = sql.Open("mysql", "root:makesense@tcp(127.0.0.1:3306)/filestore?charset=utf8")
	db.SetMaxOpenConns(1000)
	err := db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to mysql,err:" + err.Error())
		os.Exit(1)
	}
	res, err := db.Exec("select * from tbl_user")
	fmt.Println(res)

}
