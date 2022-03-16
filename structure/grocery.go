package structure

type GroceryChecklist struct {
	ID       int `json:"id"`
	Item     string `json:"item"`
	Status   int   `json:"status"`
	Quantity int    `json:"quantity"`
}