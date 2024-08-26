package database

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/mwives/go-expert/apis/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupProductDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	err = db.AutoMigrate(&entity.Product{})
	assert.NoError(t, err)
	return db
}

func TestProduct_Create(t *testing.T) {
	t.Parallel()
	db := setupProductDB(t)
	product, err := entity.NewProduct("Product 1", 10.00)
	assert.NoError(t, err)

	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)
}

func TestProduct_FindAll(t *testing.T) {
	t.Parallel()
	db := setupProductDB(t)

	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100.0)
		assert.NoError(t, err)
		db.Create(product)
	}

	productDB := NewProduct(db)

	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 23", products[2].Name)
}

func TestProduct_FindById(t *testing.T) {
	t.Parallel()
	db := setupProductDB(t)

	product, err := entity.NewProduct("Product 1", 10.0)
	assert.NoError(t, err)
	db.Create(product)
	productDB := NewProduct(db)

	product, err = productDB.FindById(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Product 1", product.Name)
}

func TestProduct_Update(t *testing.T) {
	t.Parallel()
	db := setupProductDB(t)

	product, err := entity.NewProduct("Product 1", 10.0)
	assert.NoError(t, err)
	db.Create(product)
	productDB := NewProduct(db)

	product.Name = "Product 2"
	err = productDB.Update(product)
	assert.NoError(t, err)

	updatedProduct, err := productDB.FindById(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Product 2", updatedProduct.Name)
}

func TestProduct_Delete(t *testing.T) {
	t.Parallel()
	db := setupProductDB(t)

	product, err := entity.NewProduct("Product 1", 10.0)
	assert.NoError(t, err)
	db.Create(product)
	productDB := NewProduct(db)

	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)

	deletedProduct, err := productDB.FindById(product.ID.String())
	assert.Error(t, err)
	assert.Nil(t, deletedProduct)
}
