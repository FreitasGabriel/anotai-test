package request

type CategoryRequest struct {
	Title       string `json:"title" binding:"required,min=4,max=200"`
	Description string `json:"description" binding:"required,min=4,max=200"`
	OwnerID     string `json:"owner_id" binding:"required,min=1,max=50"`
}
