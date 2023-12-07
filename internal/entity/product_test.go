package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	name := "Product 1"
	price := 10.10
	p, err := NewProduct(name, price)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, name, p.Name)
	assert.Equal(t, price, p.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", 10)
	assert.Nil(t, p)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct("test", 0)
	assert.Nil(t, p)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	p, err := NewProduct("test", -10)
	assert.Nil(t, p)
	assert.Equal(t, ErrInvalidPrice, err)
}

func TestProductValide(t *testing.T) {
	p, err := NewProduct("test 1", 10)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}
