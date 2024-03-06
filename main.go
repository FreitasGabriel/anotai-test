package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/FreitasGabriel/anotai-test/src/handler/health"
	"github.com/FreitasGabriel/anotai-test/src/handler/product"
	dbConnection "github.com/FreitasGabriel/anotai-test/src/repository"
	repo "github.com/FreitasGabriel/anotai-test/src/repository/product"
	"github.com/bemobi/log"
	"github.com/gorilla/mux"
)

func main() {

	ctx := context.Background()
	r := mux.NewRouter()
	Logger := &log.Context{Emitter: log.New(os.Stdout, time.RFC3339), Tag: "anotai-test"}
	database := repo.ProductDB{Logger: Logger, DB: dbConnection.GetConnectionDB(ctx, Logger), Context: &ctx}
	productHandler := product.ProductHandler{Logger: Logger, DB: database}

	r.HandleFunc("/product", productHandler.CreateProductHandler).Methods("POST")
	r.HandleFunc("/product/{id:[0-9]+}", productHandler.GetProductHandler).Methods("GET")
	r.HandleFunc("/product/{id:[0-9]+}", productHandler.UpdateProductHandler).Methods("PUT")
	r.HandleFunc("/product/list", productHandler.ListProductHandler).Methods("GET")

	r.HandleFunc("/health", health.HealthCheck).Methods("GET")

	http.Handle("/", r)
	server := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	Logger.I("Servidor rodando na porta")

	err := server.ListenAndServe()
	if err != nil {
		Logger.E("Erro ao iniciar o servidor")
	}

}
