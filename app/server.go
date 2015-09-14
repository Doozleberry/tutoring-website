package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/site/util"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "postgres"
	HOST        = "localhost"
	PORT        = "5432"
)

var (
	dbinfo = fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME, HOST, PORT)

	db = createDBConnection()
)

func main() {
	Run()
}

func createDBConnection() *sql.DB {
	db, err := sql.Open("postgres", dbinfo)
	util.CheckErr(err)
	util.CheckErr(db.Ping())
	return db
}
