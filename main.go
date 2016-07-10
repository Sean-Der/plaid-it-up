package main

import (
	"log"

	"github.com/sean-der/plaid-it-up/db"
)

func main() {
	err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.CreateCustomer("Jane Woods")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.CreateCustomer("Michael Li")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.CreateCustomer("Heidi Hasselbach")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.CreateCustomer("Rahul Pour")
	if err != nil {
		log.Fatal(err)
	}

}
