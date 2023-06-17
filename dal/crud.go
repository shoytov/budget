package dal

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

type Bill struct {
	id      int
	Name    string
	Balance int
}

var sqliteErr sqlite3.Error

func AddBill(newBill Bill) string {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("insert into bills (name, balance) values ($1, $2)",
		newBill.Name, newBill.Balance)
	if err != nil {
		if errors.As(err, &sqliteErr) {
			return "The bill with this name already exists!"
		}
		panic(err)
	}
	return ""
}

func GetBills() []Bill {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from bills;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	bills := []Bill{}

	for rows.Next() {
		item := Bill{}
		err := rows.Scan(&item.id, &item.Name, &item.Balance)
		if err != nil {
			fmt.Println(err)
			continue
		}
		bills = append(bills, item)
	}

	return bills
}
