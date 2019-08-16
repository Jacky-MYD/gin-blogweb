package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var SqlDB * sql.DB
func InitMysql()  {
	var err error
	SqlDB, err = sql.Open("mysql", "root:123456@(localhost:3306)/blogTest?parseTime=true")
	// 错误检查
	if err != nil {
		log.Fatal(err.Error())
	}
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

}

// 查询
func QueryRowDB(sql string) *sql.Row{
	return SqlDB.QueryRow(sql)
}

func QueryDB(sql string) (*sql.Rows, error) {
	return SqlDB.Query(sql)
}