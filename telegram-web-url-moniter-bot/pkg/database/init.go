package database

import (
	//"database/sql"
	"github.com/jmoiron/sqlx"
)

func initDatabase(db *sqlx.DB) {
	CreateTable(db, "CREATE TABLE url ( rid INTEGER PRIMARY KEY AUTOINCREMENT, url VARCHAR(256),  create_timestamp INTEGER, last_timestamp INTEGER, uid INTEGER NOT NULL DEFAULT '0', last_statuscode INTEGER NOT NULL DEFAULT '200', last_status TEXT NOT NULL DEFAULT '200 OK');")
	CreateTable(db, "CREATE TABLE user ( uid INTEGER PRIMARY KEY AUTOINCREMENT, telegram_id INTEGER, start_timestamp INTEGER, last_timestamp INTEGER);")
}