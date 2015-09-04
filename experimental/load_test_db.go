package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
   "time"
)

const (
    DB_USER     = "postgres"
    DB_PASSWORD = "postgres"
    DB_NAME		= "postgres"
    HOST		= "localhost"
    PORT		= "5432"
)

func main() {
    dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
        DB_USER, DB_PASSWORD, DB_NAME, HOST, PORT)
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    
    defer db.Close()
    
	dbvalue := "no"
	if db!=nil {
		dbvalue = "yes"
	}
	fmt.Println("# db value " + dbvalue)
	
    fmt.Println("# Inserting values")

    var lastInsertId string
    err = db.QueryRow("INSERT INTO users(username,password,last_name, first_name) VALUES($1,$2,$3, $4) returning username;", "testuser", "testpostgres", "John", "Doe").Scan(&lastInsertId)
    checkErr(err)
    fmt.Println("last inserted id =", lastInsertId)

    fmt.Println("# Updating")
    stmt, err := db.Prepare("update users set username=$1 where password=$2")
    checkErr(err)

    res, err := stmt.Exec("testuser", lastInsertId)
    checkErr(err)

    affect, err := res.RowsAffected()
    checkErr(err)

    fmt.Println(affect, "rows changed")

    fmt.Println("# Querying")
    rows, err := db.Query("SELECT * FROM users")
    checkErr(err)

    for rows.Next() {
        var password string
        var username string
        var last_name string
        var first_name string
        var created time.Time
        err = rows.Scan(&password, &username, &last_name, &first_name)
        checkErr(err)
        fmt.Println("password | username | last_name | first_name | created")
        fmt.Printf("%3v | %8v | %6v | %6v | %6v\n", password, username, last_name, first_name, created)
    }

    fmt.Println("# Deleting")
    stmt, err = db.Prepare("delete from users where password=$1")
    checkErr(err)

    res, err = stmt.Exec(lastInsertId)
    checkErr(err)

    affect, err = res.RowsAffected()
    checkErr(err)

    fmt.Println(affect, "rows changed")
}

 func checkErr(err error) {
    if err != nil {
        panic(err)
    }


}
