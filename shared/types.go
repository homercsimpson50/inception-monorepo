package shared

// Item represents a catalog item.
type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"` // cents
}
