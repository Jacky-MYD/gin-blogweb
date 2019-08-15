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


//--------图片--------
//func CreateTableWithAlbum() {
//	sql := `create table if not exists album(
//        id int(4) primary key auto_increment not null,
//        filepath varchar(255),
//        filename varchar(64),
//        status int(4),
//        createtime int(10)
//        );`
//	utils.ModifyDB(sql)
//}

