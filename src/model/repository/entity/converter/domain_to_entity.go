package converter

import (
	domain "github.com/FreitasGabriel/anotai-test/src/model/domain/category"
	productDomain "github.com/FreitasGabriel/anotai-test/src/model/domain/product"
	"github.com/FreitasGabriel/anotai-test/src/model/repository/entity"
)

func ProductDomainToEntity(pd productDomain.ProductDomainInterface) *entity.ProductEntity {
	return &entity.ProductEntity{
		ID:          pd.GetID(),
		Title:       pd.GetTitle(),
		Description: pd.GetDescription(),
		Price:       pd.GetPrice(),
		CategoryID:  pd.GetCategoryID(),
		OwnerID:     pd.GetOwnerID(),
	}
}

func CategoryDomainToEntity(cd domain.CategoryDomainInterface) *entity.CategoryEntity {
	return &entity.CategoryEntity{
		ID:          cd.GetID(),
		Title:       cd.GetTitle(),
		Description: cd.GetDescription(),
		OwnerID:     cd.GetOwnerID(),
	}
}
