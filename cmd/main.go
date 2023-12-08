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

	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products/{id}", productHandler.GetByIDProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)

	fmt.Println("Server is listening on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}

}
