package main

import (
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func main() {

	dsn := "host=localhost user=postgres password=postgres dbname=testdb port=5432 sslmode=disable"

	var err error

	db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/users/", userHandler)

	log.Println("Server started on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}