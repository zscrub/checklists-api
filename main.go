package main

import (
	"log"

	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

func main() {
	var err error
	db, err = init_db()
	if err != nil {
		log.Println(err)
		return
	}
	run_routes()
}
