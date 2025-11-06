package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Notebook", 5000)

	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.Equal(t, "Notebook", product.Name)
	assert.Equal(t, 5000, product.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	product, err := NewProduct("", 5000)

	assert.Nil(t, product)
	assert.Equal(t, ErrNameRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	product, err := NewProduct("Notebook", 0)

	assert.Nil(t, product)
	assert.Equal(t, ErrPriceRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	product, err := NewProduct("Notebook", -5000)

	assert.Nil(t, product)
	assert.Equal(t, ErrInvalidPrice, err)
}

func TestProductValidate(t *testing.T) {
	product, err := NewProduct("Notebook", 5000)

	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Nil(t, product.Validate(), err)
}
