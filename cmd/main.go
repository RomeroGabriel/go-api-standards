package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	dbModule "github.com/RomeroGabriel/go-api-standards/internal/infra/db"
	"github.com/RomeroGabriel/go-api-standards/internal/infra/webserver/handlers"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	productDB := dbModule.NewProductDB(db)
	productHandler := handlers.NewProductHandler(productDB)
	http.HandleFunc("/", productHandler.CreateProduct)
	fmt.Println("Server is listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
