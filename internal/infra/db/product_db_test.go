package db

import (
	"testing"

	"github.com/RomeroGabriel/go-api-standards/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	db := createDb()
	pDB := NewProductDB(db)
	user, _ := entity.NewProduct("SSD", 100)
	err := pDB.Create(user)
	assert.Nil(t, err)
	deleteDbTest()
}

func TestProductFindAll(t *testing.T) {
	db := createDb()
	pDB := NewProductDB(db)
	user, _ := entity.NewProduct("SSD", 100)
	_ = pDB.Create(user)
	user2, _ := entity.NewProduct("SSD2", 200)
	_ = pDB.Create(user2)
	user3, _ := entity.NewProduct("SSD3", 200)
	_ = pDB.Create(user3)

	products, err := pDB.FindAll(1, 15, "")
	assert.Nil(t, err)
	assert.NotEmpty(t, products)
	assert.Equal(t, len(products), 3)
	deleteDbTest()
}
