package model

import (
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"
	// register mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

//InitMysql init mysql connection
//2019/08/08 11:09:16
func InitMysql(dsn string) {
	if dsn == "" {
		panic("mysql dsn is empty")
	}
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("Open mysql error: %v", err))
	}

	if err = db.Ping(); err != nil {
		panic(fmt.Sprintf("Ping mysql error: %v", err))
	}

	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)
	logrus.Info("Connected to mysql.")
}
