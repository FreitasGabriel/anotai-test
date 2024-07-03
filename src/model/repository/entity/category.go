package entity

type CategoryEntity struct {
	ID          string `bson:"id,omitempty"`
	Title       string `bson:"title"`
	Description string `bson:"description"`
	OwnerID     string `bson:"owner_id"`
}
