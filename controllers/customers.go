package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sonnyochoa/CRM-Backend/models"
)

// GetHome serves the home page from a static file.
func GetHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, "./templates/home.gohtml")
}

// getCustomers returns all customers in JSON format.
func GetCustomers(customers map[string]models.Customer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customers)
	}
}

// getCustomer returns a customer object in JSON format from the passed in user ID.
// If the user ID is not found, it returns nil and StatusNotFound.
func GetCustomer(customers map[string]models.Customer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id := mux.Vars(r)["id"]

		if _, ok := customers[id]; ok {
			w.WriteHeader(http.StatusFound)
			json.NewEncoder(w).Encode(customers[id])
		} else {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(nil)
		}
	}
}

// // addCustomer creates a new customer with a unique ID based on the user data submitted.
// // If it's unable to umarshal the post request, it logs an error and returns
// // StatusBadRequest.
func AddCustomer(customers map[string]models.Customer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		uuid := uuid.New()
		id := uuid.ID()

		var customer models.Customer

		strId := strconv.FormatUint(uint64(id), 10)
		customer.ID = id

		body, _ := ioutil.ReadAll(r.Body)
		if err := json.Unmarshal(body, &customer); err != nil {
			http.Error(w, "Cannot unmarshal post request", http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(nil)
		} else {
			customers[strId] = customer

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(customers[strId])
		}
	}
}

// // updateCustomer updates the customer based on a user ID. It receives a JSON object
// // with the fields that need updating. It returns nil and StatusNotFound if the user ID
// // does not exist.
func UpdateCustomer(customers map[string]models.Customer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var customer models.Customer
		userId := mux.Vars(r)["id"]

		if _, ok := customers[userId]; ok {
			body, _ := ioutil.ReadAll(r.Body)
			if err := json.Unmarshal(body, &customer); err != nil {
				http.Error(w, "Cannot unmarshal put request", http.StatusNotFound)
			} else {
				customerID, _ := strconv.ParseUint(userId, 10, 32)
				customer.ID = uint32(customerID)
				customers[userId] = customer
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(customer)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(nil)
		}
	}
}

// // deleteCustomer deletes a customer based on a user ID.  It returns nil and
// // StatusNotFound if the user ID does not exist.
func DeleteCustomer(customers map[string]models.Customer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		userId := mux.Vars(r)["id"]

		if _, ok := customers[userId]; ok {
			delete(customers, userId)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(customers)
		} else {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(nil)
		}
	}
}
