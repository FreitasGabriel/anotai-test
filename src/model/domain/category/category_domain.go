package domain

type categoryDomain struct {
	id          string
	title       string
	description string
	ownerID     string
}

func (cd *categoryDomain) SetID(id string)        { cd.id = id }
func (cd *categoryDomain) GetID() string          { return cd.id }
func (cd *categoryDomain) GetTitle() string       { return cd.title }
func (cd *categoryDomain) GetDescription() string { return cd.description }
func (cd *categoryDomain) GetOwnerID() string     { return cd.ownerID }
