package request

type ProductRequest struct {
	Title       string `json:"title" binding:"required,min=4,max=200"`
	CategoryID  string `json:"category_id" binding:"required"`
	Price       int    `json:"price" binding:"required,min=2,max=20000"`
	Description string `json:"description" binding:"required,min=4,max=200"`
	OwnerID     string `json:"owner_id" binding:"required,min=1,max=50"`
}
