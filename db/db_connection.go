package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	// Open ini adalah database pooling yaitu management koneksi database
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/abdu")
	if err != nil {
		panic(err)
	}

	// setting database/connection pooling
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
