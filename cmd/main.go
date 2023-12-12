package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/RomeroGabriel/go-api-standards/configs"
	dbModule "github.com/RomeroGabriel/go-api-standards/internal/infra/db"
	"github.com/RomeroGabriel/go-api-standards/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	productDB := dbModule.NewProductDB(db)
	productHandler := handlers.NewProductHandler(productDB)
	userDb := dbModule.NewUserDB(db)
	userHandler := handlers.NewUserHandler(userDb)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetByIDProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
		r.Get("/", productHandler.GetProducts)
	})

	r.Route("/users", func(r chi.Router) {
		r.Use(middleware.WithValue("jwt", config.TokenAuth))
		r.Use(middleware.WithValue("JWTExperesIn", config.JWTExperesIn))
		r.Post("/", userHandler.CreateUser)
		r.Post("/token", userHandler.GetJWT)
	})

	fmt.Println("Server is listening on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}

}
