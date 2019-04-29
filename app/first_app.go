package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fun()
}

func fun() {
	db, err := sql.Open("sqlite3", "/Users/shubham/sqlite-db/gramya/data.db")
	if err != nil {
		fmt.Print(err)
		return
	}

	result, err := db.Query("select * from person")
	if err != nil {
		panic("No data from db")
	}
	for result.Next() {
	}
	fmt.Print(result)
}

func insert(db sql.DB) {
	fmt.Printf("Hello\n")
	st, err := db.Prepare("INSERT INTO PERSON (id, f_name, l_name, dob, gender) VALUES (?, ?, ?, ?, ?)")
	if err == nil {
		result, err := st.Exec(1, "Karan", "Arjun", "1700-01-01", "M")
		if err != nil {
			fmt.Print(err)
		} else {
			fmt.Print(result)
		}
	}
	fmt.Printf("\nDone")
}
