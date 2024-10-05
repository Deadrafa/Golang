package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func test() {
	connStr := "user=postgres  dbname=testdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, age int
		var name string
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			panic(err)
		}
		fmt.Printf("ID: %d, Username %s, Age: %d\n", id, name, age)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

}
