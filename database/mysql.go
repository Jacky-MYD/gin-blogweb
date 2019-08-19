package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var SqlDB * sql.DB
func InitMysql() (err error) {
	fmt.Println("2324")
	SqlDB, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/blogTest?parseTime=true")
	// 错误检查
	if err != nil {
		log.Fatal(err.Error())
	}
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	return

}

// 查询
func QueryRowDB(sql string) *sql.Row{
	return SqlDB.QueryRow(sql)
}

func QueryDB(sql string) (*sql.Rows, error) {
	return SqlDB.Query(sql)
}