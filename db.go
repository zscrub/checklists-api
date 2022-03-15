package main

import (
	"fmt"
	"zscrub/checklists-api/structure"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func init_db() ([]structure.GroceryChecklist, error) {
	db, err := sqlx.Open("sqlite3", "data/checklists.db")
	if err != nil {
		return nil, err
	}

	var data []structure.GroceryChecklist
	sql_get_groceries := "select * from groceries;"

	err = db.Select(&data, sql_get_groceries)
	if err != nil {
		return nil, err
	}

	fmt.Println(data)

	return data, nil
}
