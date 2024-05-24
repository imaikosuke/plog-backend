package db

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq"
)

var db *sql.DB

func Init() {
    var err error
    connStr := "user=imaikosuke password=postgresql0202 dbname=plog sslmode=disable"
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    if err = db.Ping(); err != nil {
        log.Fatal(err)
    }

    fmt.Println("Successfully connected to the database")
}

func GetDB() *sql.DB {
    return db
}
