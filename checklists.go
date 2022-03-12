package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"zscrub/checklists-api/structure"
)



func read_grocery_data() []structure.Grocery_checklist {
	content, _ := ioutil.ReadFile("./data.json")

	var data []structure.Grocery_checklist
	err := json.Unmarshal(content, &data)

	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	return data
}
