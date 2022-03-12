package structure

type Movie_checklist struct {
	ID     int    `json:"id"`
	Movie  string `json:"movie"`
	Status bool   `json:"status"`
}