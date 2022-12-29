package db_test

import (
	"database/sql"
	"github.com/iamp5/go-hexagonal/adapters/db"
	"github.com/iamp5/go-hexagonal/application"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

var Db *sql.DB

func setup() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
			"id" string,
			"name" string,
			"status" string,
			"price" float
			);`

	stmt, err := db.Prepare(table)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products VALUES (
			"abc",
			"Product Test",
			"disabled",
			0
			);`

	stmt, err := db.Prepare(insert)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setup()
	defer Db.Close()

	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("abc")

	require.Nil(t, err)
	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t, "disabled", product.GetStatus())
	require.Equal(t, 0.0, product.GetPrice())
}

func TestProductDb_Save(t *testing.T) {
	setup()
	defer Db.Close()
	productDb := db.NewProductDb(Db)

	product := application.NewProduct()
	product.Name = "Product Test"
	product.Price = 25

	productResult, err := productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Status, productResult.GetStatus())
	require.Equal(t, product.Price, productResult.GetPrice())

	product.Status = "enabled"
	productResult, err = productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Status, productResult.GetStatus())
	require.Equal(t, product.Price, productResult.GetPrice())
}
