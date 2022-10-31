package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/google/uuid"
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

func getHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, "./templates/home.gohtml")
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customers)
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

func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicaiton/json")

	uuid := uuid.New()
	id := uuid.ID()

	var customer models.Customer

	strId := strconv.FormatUint(uint64(id), 10)
	customer.ID = id

	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, &customer); err != nil {
		fmt.Println("Cannot unmarshal post request")
	}
	customers[strId] = customer

	json.NewEncoder(w).Encode(customers[strId])
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userId := mux.Vars(r)["id"]

	if _, ok := customers[userId]; ok {
		delete(customers, userId)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customers)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(customers)
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", getHome).Methods("GET")
	r.HandleFunc("/customers", getCustomers).Methods("GET")
	r.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	r.HandleFunc("/customers", addCustomer).Methods("POST")
	r.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")
	fmt.Println("Starting server on localhost:3000...")
	http.ListenAndServe(":3000", r)
}
