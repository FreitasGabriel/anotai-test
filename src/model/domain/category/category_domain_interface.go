package domain

type CategoryDomainInterface interface {
	SetID(id string)
	GetID() string
	GetTitle() string
	GetDescription() string
	GetOwnerID() string
}

func NewCategoryDomain(id, title, description, ownerID string) CategoryDomainInterface {
	return &categoryDomain{
		id:          id,
		title:       title,
		description: description,
		ownerID:     ownerID,
	}
}
