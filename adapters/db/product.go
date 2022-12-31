package db

import (
	"database/sql"
	"github.com/iamp5/go-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db}
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {
	var product application.Product

	stmt, err := p.db.Prepare("select id, name, status, price from products where id = ?")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Status, &product.Price)

	return &product, nil
}

func (p *ProductDb) Save(product application.ProductInterface) (application.ProductInterface, error) {
	var id string
	err := p.db.QueryRow("select id from products where id = ?", product.GetId()).Scan(&id)

	if err != nil {
		_, err = p.create(product)
		if err != nil {
			return nil, err
		}

		return product, nil
	}

	_, err = p.update(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

// private function
func (p *ProductDb) create(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := p.db.Prepare(`insert into products(id, name, status, price) values (?, ?, ?, ?)`)
	defer stmt.Close()

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(product.GetId(), product.GetName(), product.GetStatus(), product.GetPrice())
	if err != nil {
		return nil, err
	}

	return product, nil
}

// private function
func (p *ProductDb) update(product application.ProductInterface) (application.ProductInterface, error) {
	_, err := p.db.Exec(`update products set name = ?, status = ?, price = ? where id = ?`,
		product.GetName(), product.GetStatus(), product.GetPrice(), product.GetId())

	if err != nil {
		return nil, err
	}

	return product, nil
}
