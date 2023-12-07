package db

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/RomeroGabriel/go-api-standards/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	db := createDb()
	pDB := NewProductDB(db)
	product, err := entity.NewProduct("SSD", 100)
	assert.NoError(t, err)
	err = pDB.Create(product)
	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)
	deleteDbTest()
}

func TestProductFindAll(t *testing.T) {
	db := createDb()
	pDB := NewProductDB(db)

	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64())
		assert.NoError(t, err)
		err = pDB.Create(product)
		assert.NoError(t, err)
	}

	products, err := pDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = pDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)
	deleteDbTest()
}

func TestProductFindByID(t *testing.T) {
	db := createDb()
	pDB := NewProductDB(db)

	product, err := entity.NewProduct("SSD", 100)
	assert.NoError(t, err)
	pDB.Create(product)

	productFind, err := pDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "SSD", productFind.Name)

	deleteDbTest()
}

func TestProductUpdate(t *testing.T) {
	db := createDb()
	pDB := NewProductDB(db)

	product, err := entity.NewProduct("SSD", 100)
	assert.NoError(t, err)
	pDB.Create(product)

	product.Name = "Test Update"
	err = pDB.Update(product)
	assert.NoError(t, err)
	productFind, err := pDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Test Update", productFind.Name)

	deleteDbTest()
}

func TestProductDelete(t *testing.T) {
	db := createDb()
	pDB := NewProductDB(db)

	product, err := entity.NewProduct("SSD", 100)
	assert.NoError(t, err)
	pDB.Create(product)

	err = pDB.Delete(product)
	assert.NoError(t, err)
	productFind, err := pDB.FindByID(product.ID.String())
	assert.Error(t, err)
	assert.Nil(t, productFind)

	deleteDbTest()
}
