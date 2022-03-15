package structure

type GroceryChecklist struct {
	ID       int `json:"id"`
	Item     string `json:"item"`
	Status   bool   `json:"status"`
	Quantity int    `json:"quantity"`
}