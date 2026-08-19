package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func main(){
	dsn := "host=localhost user=postgres password=postgres dbname=orderdb port=5432 sslmode=disable"

	var err error
	
	db, err = sqlx.Connect("postgres", dsn)

	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	router.HandleFunc("GET /orders", GetOrders)
	router.HandleFunc("POST /orders", CreateOrder)
	router.HandleFunc("PUT /orders", UpdateOrder)
	router.HandleFunc("GET /orders/{id}", GetOrder)
	router.HandleFunc("DELETE /orders/{id}", DeleteOrder)

	http.Handle("/", router)
	http.ListenAndServe(":8081", nil)
}