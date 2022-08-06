package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// db, err := sql.Open("sqlserver", "sqlserver://sa:P@ssw0rd@13.76.163.73?datbase=techcoach")
	db, err := sql.Open("mysql", "addmin:0988327674Mysql@tcp(localhost:3306)/sys")
	if err != nil {
		panic(err)
	} else {
		println("No error on connect db")
	}

	err = db.Ping()
	if err != nil {
		panic(err)
		// println("can not connect the local database")
	}

	println("can connect the local database")
	

	// query := "select * from cover"
	// rows, err := db.Query(query)
	// if err != nil {
	// 	panic(err)
	// }

	// id := 0
	// name := ""
	// ok := rows.Next()
	// if ok {
	// 	rows.Scan(&id, &name)
	// }
	// println(id, name)
}
