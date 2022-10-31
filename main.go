package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	http.ServeFile(w, r, "./templates/home.gohtml")
}

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world\n")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", GetHome).Methods("GET")
	r.HandleFunc("/customers", GetCustomers).Methods("GET")
	fmt.Println("Starting server on localhost:3000...")
	http.ListenAndServe(":3000", r)
}
