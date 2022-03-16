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

// gets all groceries
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

// in one API call, loop over multiple groceries to insert statement/execute n queries
func go_post_groceries(groceries []structure.GroceryChecklist) (int, error) {

	for _, obj := range groceries {
		result, err := db.Exec("insert into groceries (Item, Status, Quantity) VALUES (?, ?, ?)", obj.Item, obj.Status, obj.Quantity)

		// need to use proper response code
		if err != nil {
			return 418, err
		}
		fmt.Println(result)
	}
	return 201, nil
}
