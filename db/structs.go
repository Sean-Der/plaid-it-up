package db

var schema = `
PRAGMA foreign_keys = ON;

CREATE TABLE customer (
    id   INTEGER NOT NULL,
    name TEXT    NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE account (
    id          INTEGER NOT NULL,
	balance     INTEGER NOT NULL,
	customer_id INTEGER NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY(customer_id) REFERENCES customer(id)
);

CREATE TABLE transfer (
    id              INTEGER NOT NULL,
    amount          INTEGER NOT NULL,
    src_account_id  INTEGER NOT NULL,
    dst_account_id  INTEGER NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY(src_account_id) REFERENCES account(id),
    FOREIGN KEY(dst_account_id) REFERENCES account(id)
);`

type Customer struct {
	Id   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type Account struct {
	Id       int64 `db:"id"`
	Balance  int64 `db:"balance"`
	Customer int64 `db:"customer_id"`
}

type Transfer struct {
	Id         int64 `db:"id"`
	Amount     int64 `db:"amount"`
	SrcAccount int64 `db:"src_account_id"`
	DstAccount int64 `db:"dst_account_id"`
}
