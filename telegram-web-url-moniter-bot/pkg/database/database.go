package database

import (
	//"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"github.com/jmoiron/sqlx"
)

func OpenOrCreateDatabase(filepath string) (*sqlx.DB, error) {
	if !fileExists(filepath) {
		log.Printf("Creating db file: %s...\n", filepath)
		file, err := os.Create(filepath) // Create SQLite file
		if err != nil {
			return nil, err
		}
		file.Close()
		log.Printf("%s created\n", filepath)
		sqliteDatabase, err := sqlx.Open("sqlite3", filepath) // Open the created SQLite File
		if err != nil {
			return nil, err
		}
		initDatabase(sqliteDatabase)
		return sqliteDatabase, nil
	}else {
		sqliteDatabase, err := sqlx.Open("sqlite3", filepath) // Open the created SQLite File
		if err != nil {
			return nil, err
		}
		return sqliteDatabase, nil
	}
}

func RemoveDatabase(filepath string) {
	os.Remove(filepath)
}