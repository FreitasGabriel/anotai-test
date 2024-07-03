package domain

type productDomain struct {
	id          string
	title       string
	categoryID  string
	price       int
	description string
	ownerID     string
}

func (pd *productDomain) SetID(id string)        { pd.id = id }
func (pd *productDomain) GetID() string          { return pd.id }
func (pd *productDomain) GetTitle() string       { return pd.title }
func (pd *productDomain) GetPrice() int          { return pd.price }
func (pd *productDomain) GetDescription() string { return pd.description }
func (pd *productDomain) GetCategoryID() string  { return pd.categoryID }
func (pd *productDomain) GetOwnerID() string     { return pd.ownerID }
