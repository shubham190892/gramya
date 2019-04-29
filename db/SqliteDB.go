package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"gramya/utils"
)

var DB *sql.DB

func init() {
	utils.Logger.Info("Initiating connections from databases")
	db, err := sql.Open("sqlite3", "/Users/shubham/sqlite-db/gramya/data.db")
	if err != nil {
		utils.Logger.Error(err.Error())
		panic(err)
	}
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(1)
	DB = db
}
