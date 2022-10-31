package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sonnyochoa/CRM-Backend/models"
)

var customers = []models.Customer{
	{
		ID:        1,
		Name:      "sonny",
		Role:      "admin",
		Email:     "admin@sonny.com",
		Phone:     "1234567890",
		Contacted: false,
	},
	{
		ID:        2,
		Name:      "karokomako",
		Role:      "writer",
		Email:     "writer@sonny.com",
		Phone:     "5012346789",
		Contacted: false,
	},
	{
		ID:        3,
		Name:      "churikua",
		Role:      "builder",
		Email:     "builder@sonny.com",
		Phone:     "3012456789",
		Contacted: true,
	},
}

func GetHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	http.ServeFile(w, r, "./templates/home.gohtml")
}

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customers)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", GetHome).Methods("GET")
	r.HandleFunc("/customers", GetCustomers).Methods("GET")
	fmt.Println("Starting server on localhost:3000...")
	http.ListenAndServe(":3000", r)
}
