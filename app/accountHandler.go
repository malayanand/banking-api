package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/malayanand/banking/dto"
	"github.com/malayanand/banking/service"
)

type AccountHandlers struct {
	service service.AccountService
}

func (ah *AccountHandlers) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	accountId := vars["account_id"]

	var request dto.TransactionRequest

	// decode incoming request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.AccountId = accountId
		request.CustomerId = customerId

		// make transactiion
		account, appErr := ah.service.MakeTransaction(request)
		if appErr != nil {
			writeResponse(w, appErr.Code, appErr.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, account)
		}
	}
}

func (ah *AccountHandlers) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		// add customer id fetched from params to request body
		request.CustomerId = customerId
		account, appError := ah.service.NewAccount(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.Message)
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}
