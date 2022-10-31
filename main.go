package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sonnyochoa/CRM-Backend/models"
)

var customers = map[string]models.Customer{
	"1": {
		ID:        1,
		Name:      "sonny",
		Role:      "admin",
		Email:     "admin@sonny.com",
		Phone:     "1234567890",
		Contacted: false,
	},
	"2": {
		ID:        2,
		Name:      "panda",
		Role:      "writer",
		Email:     "writer@sonny.com",
		Phone:     "5012346789",
		Contacted: false,
	},
	"3": {
		ID:        3,
		Name:      "bear",
		Role:      "builder",
		Email:     "builder@sonny.com",
		Phone:     "3012456789",
		Contacted: true,
	},
}

func GetHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, "./templates/home.gohtml")
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	if _, ok := customers[id]; ok {
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(customers[id])
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("{User does not exist}")
	}
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customers)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", GetHome).Methods("GET")
	r.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	r.HandleFunc("/customers", getCustomers).Methods("GET")
	fmt.Println("Starting server on localhost:3000...")
	http.ListenAndServe(":3000", r)
}
