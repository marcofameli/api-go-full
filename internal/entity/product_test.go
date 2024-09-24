package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProduct(t *testing.T) {

	p, err := NewProduct("Product 1", 10.0)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, "Product 1", p.Name)
	assert.Equal(t, 10.0, p.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", 10)
	assert.NotNil(t, err)
	assert.Equal(t, "Name is required", err.Error())
	assert.Nil(t, p)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct("Product 1", 0)
	assert.NotNil(t, err)
	assert.Equal(t, "Price is required", err.Error())
	assert.Nil(t, p)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	p, err := NewProduct("Product 1", -10)
	assert.NotNil(t, err)
	assert.Equal(t, "Invalid price", err.Error())
	assert.Nil(t, p)
}

func TestProduct_Validate(t *testing.T) {
	p, err := NewProduct("Product 1", 10)
	assert.Nil(t, err)
	assert.Nil(t, p.Validate())
}
