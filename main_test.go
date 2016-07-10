package main

import (
	"testing"

	"github.com/sean-der/plaid-it-up/db"
)

func TestCustomerFK(t *testing.T) {
	err := db.Connect()
	if err != nil {
		t.Error(err)
	}

	id, err := db.CreateCustomer("Tester")
	if err != nil {
		t.Error(err)
	}

	_, err = db.CreateAccount(id, 0)
	if err != nil {
		t.Error(err)
	}

	_, err = db.CreateAccount(10000, 0)
	if err == nil {
		t.Error("CreateAccount allowed a invalid customer id")
	}
}

func TestAccountFK(t *testing.T) {
	err := db.Connect()
	if err != nil {
		t.Error(err)
	}

	customerId, err := db.CreateCustomer("Tester")
	if err != nil {
		t.Error(err)
	}

	account1Id, err := db.CreateAccount(customerId, 0)
	if err != nil {
		t.Error(err)
	}

	account2Id, err := db.CreateAccount(customerId, 0)
	if err != nil {
		t.Error(err)
	}

	_, err = db.CreateTransfer(account1Id, account2Id, 0)
	if err != nil {
		t.Error(err)
	}

	_, err = db.CreateTransfer(10000, 10000, 0)
	if err == nil {
		t.Error("CreateTransfer allowed an invalid acount id")
	}

}

func TestTransferBalanceChecking(t *testing.T) {
	err := db.Connect()
	if err != nil {
		t.Error(err)
	}

	customerId, err := db.CreateCustomer("Tester")
	if err != nil {
		t.Error(err)
	}

	account1Id, err := db.CreateAccount(customerId, 500)
	if err != nil {
		t.Error(err)
	}

	account2Id, err := db.CreateAccount(customerId, 0)
	if err != nil {
		t.Error(err)
	}

	_, err = db.CreateTransfer(account1Id, account2Id, 500)
	if err != nil {
		t.Error(err)
	}

	_, err = db.CreateTransfer(account1Id, account2Id, 500)
	if err == nil {
		t.Error("CreateTransfer allowed a transfer larger than an available balance")
	}
}
