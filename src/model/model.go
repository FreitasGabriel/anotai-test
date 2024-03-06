package model

type ProductPayload struct {
	OwnerID     string  `json:"owner_id"`
	ProductID   string  `json:"product_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	CategoryID  string  `json:"category"`
}
