package app

import (
	"fmt"
	"database/sql"
)

//连接mysql
func ConnectMysql()(*sql.DB,error) {
	dbDsn :=Config.DbDsn
	dbData, err := sql.Open("mysql", dbDsn)
	if err != nil {
		fmt.Print("connect db error")
		return dbData,err
	}

	return dbData,nil
}