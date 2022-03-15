package main

import (
	"fmt"
	"zscrub/checklists-api/structure"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func init_db() (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", "data/checklists.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}


func go_get_groceries() ([]structure.GroceryChecklist, error) {
	var data []structure.GroceryChecklist
	sql_get_groceries := "select * from groceries;"

	err := db.Select(&data, sql_get_groceries)
	if err != nil {
		return nil, err
	}

	fmt.Println(data)

	return data, nil
}
