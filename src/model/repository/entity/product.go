package entity

type ProductEntity struct {
	ID          string `bson:"id,omitempty"`
	Title       string `bson:"title"`
	CategoryID  string `bson:"category_id"`
	Price       int    `bson:"price"`
	Description string `bson:"description"`
	OwnerID     string `bson:"owner_id"`
}
