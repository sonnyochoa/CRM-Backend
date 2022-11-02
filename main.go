package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sonnyochoa/CRM-Backend/controllers"
	"github.com/sonnyochoa/CRM-Backend/models"
)

var customers map[string]models.Customer

func initModel() error {
	customerDB, err := ioutil.ReadFile("./models/customers.json")

	if err != nil {
		log.Fatal(err)
		return err
	}
	dec := json.NewDecoder(strings.NewReader(string(customerDB)))
	for dec.More() {
		err = dec.Decode(&customers)
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func main() {
	r := mux.NewRouter()

	err := initModel()
	if err != nil {
		log.Fatal(err)
	}
	r.HandleFunc("/", controllers.GetHome).Methods("GET")
	r.HandleFunc("/customers", controllers.GetCustomers(customers)).Methods("GET")
	r.HandleFunc("/customers/{id}", controllers.GetCustomer(customers)).Methods("GET")
	r.HandleFunc("/customers", controllers.AddCustomer(customers)).Methods("POST")
	r.HandleFunc("/customers/{id}", controllers.UpdateCustomer(customers)).Methods("PUT")
	r.HandleFunc("/customers/{id}", controllers.DeleteCustomer(customers)).Methods("DELETE")
	fmt.Println("Starting server on localhost:3000...")
	http.ListenAndServe(":3000", r)
}
