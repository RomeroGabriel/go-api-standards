package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	dbModule "github.com/RomeroGabriel/go-api-standards/internal/infra/db"
	"github.com/RomeroGabriel/go-api-standards/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	productDB := dbModule.NewProductDB(db)
	productHandler := handlers.NewProductHandler(productDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/product", productHandler.CreateProduct)
	r.Get("/product/{id}", productHandler.GetByIDProduct)
	r.Put("/product/{id}", productHandler.UpdateProduct)
	r.Delete("/product/{id}", productHandler.DeleteProduct)
	r.Get("/products", productHandler.GetProducts)

	fmt.Println("Server is listening on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}

}
