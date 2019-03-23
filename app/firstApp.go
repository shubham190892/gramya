package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"html"
	"log"
	"net/http"
)

func main() {
	server1()
}

func server1() {
	defaultHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Default, %q", html.EscapeString(r.URL.Path))
	}
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	}

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Printf("Goring to start the server")
	log.Fatal(http.ListenAndServe(":8383", nil))
}

func fun() {
	db, err := sql.Open("sqlite3", "/Users/shubham/sqlite-db/gramya/data.db")
	if err != nil {
		fmt.Print(err)
		return
	}

	result, err := db.Query("select * from person")
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
