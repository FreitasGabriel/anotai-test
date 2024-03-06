package product

import (
	"encoding/json"
	"net/http"

	"github.com/FreitasGabriel/anotai-test/src/model"
	"github.com/FreitasGabriel/anotai-test/src/repository/product"
	productService "github.com/FreitasGabriel/anotai-test/src/service/product"
	"github.com/bemobi/log"
)

type ProductHandler struct {
	Logger *log.Context
	DB     product.ProductDB
}

func (ph *ProductHandler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {

	product := &model.ProductPayload{}

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		ph.Logger.E("Error on decode product data", err)
		http.Error(w, "Error on decode product data", http.StatusInternalServerError)
	}

	errService := productService.CreateProductService(product, ph.DB)
	if errService != nil {
		ph.Logger.E("Error on create document on colletion", errService)
		http.Error(w, "Error on create document on colletion", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	ph.Logger.I("Product created successfully")
}

func (ph *ProductHandler) GetProductHandler(w http.ResponseWriter, r *http.Request) {

}

func (ph *ProductHandler) ListProductHandler(w http.ResponseWriter, r *http.Request) {

	result, err := productService.ListProductsService(ph.DB)
	if err != nil {
		ph.Logger.E("Error to return colletions list", err)
		http.Error(w, "Error to return colletions list", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		ph.Logger.E("Error to encode list of products", "err", err)
		http.Error(w, "Error to encode list of products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	ph.Logger.I("List of products succesfully")
}

func (ph *ProductHandler) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {}
