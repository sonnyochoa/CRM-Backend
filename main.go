package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world\n")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/customers", GetCustomers).Methods("GET")
	fmt.Println("Starting server on localhost:3000...")
	http.ListenAndServe(":3000", r)
}
