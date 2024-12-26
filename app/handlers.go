package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/malayanand/banking/service"
)

/*
	Name    string `json: "full_name" xml:"name"`
type Customer struct {
	City    string `json: "city" xml:"city"`
	Zipcode string `json: "zip_code" xml:"zipcode"`
}*/

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")

	customers, err := ch.service.GetAllCustomers(status)

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
}

func (ch *CustomerHandlers) getCustomers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	// customer is received as a DTO (Data transfer object)
	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
