package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sean-der/plaid-it-up/db"
)

func customerHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		idStr := r.URL.Query().Get("id")
		if len(idStr) == 0 {
			customers, err := db.GetCustomers()
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			json.NewEncoder(w).Encode(customers)
		} else {
			id, err := strconv.ParseInt(idStr, 10, 0)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}

			customer, err := db.GetCustomer(id)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			json.NewEncoder(w).Encode(customer)
		}
	case "POST":
	default:
		http.Error(w, "Unhandled HTTP Method Type", 500)
	}
}

func accountHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		idStr := r.URL.Query().Get("id")
		if len(idStr) == 0 {
			accounts, err := db.GetAccounts()
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			json.NewEncoder(w).Encode(accounts)
		} else {
			id, err := strconv.ParseInt(idStr, 10, 0)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}

			account, err := db.GetAccount(id)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			json.NewEncoder(w).Encode(account)
		}
	case "POST":
	default:
		http.Error(w, "Unhandled HTTP Method Type", 500)
	}
}

func transferHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		idStr := r.URL.Query().Get("id")
		if len(idStr) == 0 {
			transfers, err := db.GetTransfers()
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			json.NewEncoder(w).Encode(transfers)
		} else {
			id, err := strconv.ParseInt(idStr, 10, 0)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}

			transfer, err := db.GetTransfer(id)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			json.NewEncoder(w).Encode(transfer)
		}
	case "POST":
	default:
		http.Error(w, "Unhandled HTTP Method Type", 500)
	}
}

func StartServer(port string) {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/customer", customerHandler)
	myRouter.HandleFunc("/account", accountHandler)
	myRouter.HandleFunc("/transfer", transferHandler)

	log.Fatal(http.ListenAndServe(port, myRouter))
}
