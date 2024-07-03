package domain

type ProductDomainInterface interface {
	SetID(id string)
	GetID() string
	GetTitle() string
	GetDescription() string
	GetPrice() int
	GetCategoryID() string
	GetOwnerID() string
}

func NewProductDomain(id, title, description, categoryID, ownerID string, price int) ProductDomainInterface {
	return &productDomain{
		id:          id,
		title:       title,
		description: description,
		price:       price,
		categoryID:  categoryID,
		ownerID:     ownerID,
	}
}
