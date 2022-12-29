package main

import (
	"database/sql"
	dbAdapter "github.com/iamp5/go-hexagonal/adapters/db"
	"github.com/iamp5/go-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "sqlite.db")
	productDbAdapter := dbAdapter.NewProductDb(db)

	productService := application.NewProductService(productDbAdapter)
	product, _ := productService.Create("Product 1", 30)

	productService.Enable(product)
}
