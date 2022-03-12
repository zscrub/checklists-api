package structure

type Grocery_checklist struct {
	ID       int `json:"id"`
	Item     string `json:"item"`
	Status   bool   `json:"status"`
	Quantity int    `json:"quantity"`
}