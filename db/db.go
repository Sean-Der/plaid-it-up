package db

import (
	"errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB

func Connect() (err error) {
	db, err = sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		return
	}

	db.MustExec(schema)
	return
}

func CreateCustomer(name string) (id int64, err error) {
	result, err := db.NamedExec(`INSERT INTO customer(name) VALUES (:name)`,
		map[string]interface{}{
			"name": name,
		})
	if err != nil {
		return
	}
	id, err = result.LastInsertId()
	return
}

func CreateAccount(customer_id int64, balance int64) (id int64, err error) {
	result, err := db.NamedExec(`INSERT INTO account(balance, customer_id) VALUES (:balance, :customer_id)`,
		map[string]interface{}{
			"balance":     balance,
			"customer_id": customer_id,
		})
	if err != nil {
		return
	}
	id, err = result.LastInsertId()
	return
}

func CreateTransfer(src_account_id int64, dst_account_id int64, amount int64) (id int64, err error) {
	src_account, err := getAccount(src_account_id)
	if err != nil {
		return
	}

	dst_account, err := getAccount(dst_account_id)
	if err != nil {
		return
	}

	if src_account.Balance-amount < 0 {
		err = errors.New("Account has insufficent funds for transfer")
		return
	}

	src_account.Balance = src_account.Balance - amount
	dst_account.Balance = dst_account.Balance + amount

	tx := db.MustBegin()

	err = updateAccount(src_account, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = updateAccount(dst_account, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()
	return
}

func getAccount(id int64) (account *Account, err error) {
	account = &Account{}
	err = db.Get(account, "SELECT * from account WHERE id = $1", id)
	return
}

func updateAccount(account *Account, tx *sqlx.Tx) (err error) {
	_, err = tx.NamedExec(`UPDATE account set balance = :balance WHERE id = :id`,
		map[string]interface{}{
			"balance": account.Balance,
			"id":      account.Id,
		})
	return
}