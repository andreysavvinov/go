package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
}

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Неверное сообщение: ", http.StatusBadRequest)
		return
	}

	var res Response
	res.Message = "Привет, " + req.Name
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Неверное сообщение: ", http.StatusBadRequest)
		return
	}
}

func main() {
	http.HandleFunc("/", HandleFunc)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
