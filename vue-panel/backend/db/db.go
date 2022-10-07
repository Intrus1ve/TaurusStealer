package db

import (
	"log"

	config "../config"
	lggr "../logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB
var logger = lggr.Log()

func SQLConnect() *sqlx.DB {
	pass := config.GetConfig().Db_pass
	var err error
	db, err = sqlx.Open("mysql", "root:"+pass+"@tcp/taurus")
	if err != nil {
		logger.Println("db.Connect: ", err)
		log.Fatal("Connecting to DB:", err)
	}

	err = db.Ping()
	if err != nil {
		logger.Println("db.Connect: ", err)
		log.Fatal("Connecting to DB:", err)
	}

	db.SetMaxIdleConns(0)
	db.SetMaxOpenConns(500)

	return db
}

func Connect() *sqlx.DB {
	if db != nil {
		err := db.Ping()
		if err == nil {
			return db
		}
	}
	return SQLConnect()
}
