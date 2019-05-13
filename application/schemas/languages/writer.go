package languages

type Writer struct {
	Title       string `json:"title"`
	NewOrder    string `json:"new_order"`
	OrderNumber string `json:"order_number"`
	OrderType   string `json:"order_type"`
}
