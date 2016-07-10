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
		var customer db.Customer
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&customer); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		id, err := db.CreateCustomer(customer.Name)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		createdCustomer, err := db.GetCustomer(id)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		json.NewEncoder(w).Encode(createdCustomer)
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
		var account db.Account
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&account); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		id, err := db.CreateAccount(account.Customer, account.Balance)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		createdAccount, err := db.GetAccount(id)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		json.NewEncoder(w).Encode(createdAccount)
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
		var transfer db.Transfer
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&transfer); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		id, err := db.CreateTransfer(transfer.SrcAccount, transfer.DstAccount, transfer.Amount)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		createdTransfer, err := db.GetTransfer(id)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		json.NewEncoder(w).Encode(createdTransfer)
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
