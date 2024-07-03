package converter

import (
	domain "github.com/FreitasGabriel/anotai-test/src/model/domain/category"
	productDomain "github.com/FreitasGabriel/anotai-test/src/model/domain/product"
	"github.com/FreitasGabriel/anotai-test/src/model/repository/entity"
)

func ProductEntityToDomain(entity entity.ProductEntity) productDomain.ProductDomainInterface {
	productDomain := productDomain.NewProductDomain(
		entity.ID,
		entity.Title,
		entity.Description,
		entity.CategoryID,
		entity.OwnerID,
		entity.Price,
	)
	productDomain.SetID(entity.ID)
	return productDomain
}

func CategoryEntityToDomain(entity entity.CategoryEntity) domain.CategoryDomainInterface {
	categoryDomain := domain.NewCategoryDomain(
		entity.ID,
		entity.Title,
		entity.Description,
		entity.OwnerID,
	)

	categoryDomain.SetID(entity.ID)
	return categoryDomain
}
