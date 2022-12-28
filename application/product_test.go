package application_test

import (
	"github.com/iamp5/go-hexagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{
		ID:     "1",
		Name:   "Product 1",
		Price:  10,
		Status: application.DISABLED,
	}

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{
		ID:     "1",
		Name:   "Product 1",
		Price:  0,
		Status: application.ENABLED,
	}

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	product.Enable()

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "the price must be equals to zero to disable the product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{
		ID:     uuid.NewV4().String(),
		Name:   "Product 1",
		Price:  10,
		Status: application.DISABLED,
	}

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "invalid"
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater or equal to zero", err.Error())
}
