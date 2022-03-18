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
func go_post_groceries(groceries []structure.GroceryChecklist) error {

	query_insert := "insert into groceries (Item, Status, Quantity) VALUES "
	params := []interface{}{}
	for i, obj := range groceries {
		p1 := i * 3

		query_insert += fmt.Sprintf("($%d,$%d,$%d),", p1, p1+1, p1+2)

		params = append(params, obj.Item, obj.Status, obj.Quantity)
	}
	query_insert = query_insert[:len(query_insert)-1]

	fmt.Println(query_insert)
	fmt.Println(params)

	result, err := db.Exec(query_insert, params...)
	if err != nil {
		return err
	}

	fmt.Println(result)

	return nil
}

func go_delete_groceries(id_arr []int) error {

	query_delete := "delete from groceries where id in ("

	for _, id := range id_arr {
		query_delete += fmt.Sprintf("%d,",id)
	}
	query_delete = query_delete[:len(query_delete)-1]
	query_delete += ")"

	result, err := db.Exec(query_delete)
	if err != nil {
		return err
	}

	fmt.Println(result)

	return nil
}
