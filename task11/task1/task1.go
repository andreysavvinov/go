package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "Привет, %s!\n", name)
	w.Write(body)
}

func main() {
	http.HandleFunc("/echo", HandleFunc)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
