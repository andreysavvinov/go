package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func getUsers(w http.ResponseWriter, r *http.Request) {

	var users []User

	err := db.Select(&users, "SELECT * FROM users ORDER BY id")

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = db.QueryRow(
		`
		INSERT INTO users(name,email,age,is_active)
		VALUES($1,$2,$3,$4)
		RETURNING id
		`,
		user.Name,
		user.Email,
		user.Age,
		user.IsActive,
	).Scan(&user.ID)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	_, err = db.Exec(`
		UPDATE users
		SET
			name=$1,
			email=$2,
			age=$3,
			is_active=$4
		WHERE id=$5
	`,
		user.Name,
		user.Email,
		user.Age,
		user.IsActive,
		user.ID,
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte("updated"))
}

func deleteUser(w http.ResponseWriter, r *http.Request) {

	idStr := strings.TrimPrefix(r.URL.Path, "/users/")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "invalid id", 400)
		return
	}

	_, err = db.Exec(
		"DELETE FROM users WHERE id=$1",
		id,
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte("deleted"))
}

func getUser(w http.ResponseWriter, r *http.Request) {

	idStr := strings.TrimPrefix(r.URL.Path, "/users/")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "invalid id", 400)
		return
	}

	var user User

	err = db.Get(
		&user,
		"SELECT * FROM users WHERE id=$1",
		id,
	)

	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		getUsers(w, r)

	case http.MethodPost:
		createUser(w, r)

	case http.MethodPut:
		updateUser(w, r)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		getUser(w, r)

	case http.MethodDelete:
		deleteUser(w, r)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}