package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
	var orders []Order

	err := db.Select(&orders, "SELECT * FROM orders ORDER BY id")

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order Order

	err := json.NewDecoder(r.Body).Decode(&order)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = db.QueryRow(`INSERT INTO orders(user_id, product, price) VALUES ($1, $2, $3) RETURNING id`,
		order.UserID,
		order.Product,
		order.Price).Scan(&order.ID)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(order)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	var order Order

	err := json.NewDecoder(r.Body).Decode(&order)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	_, err = db.Exec(`UPDATE orders
		user_id = $1,
		product = $2,
		price = $3
		WHERE id = $4`, order.UserID,
		order.Product,
		order.Price,
		order.ID)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte("updated"))
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	var order Order

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "invalid id", 400)
		return
	}

	err = db.Get(&order, `SELECT * FROM orders WHERE id=$1`, id)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(order)
}

func DeleteOrder(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "invalid id", 400)
		return
	}

	_, err = db.Exec(`DELETE FROM orders WHERE id=$1`, id)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte("deleted"))
}
