package entity

import (
	"testing"

	"github.com/mwives/go-expert/apis/pkg/entity"
	"github.com/stretchr/testify/assert"
)

func TestProduct_NewProduct(t *testing.T) {
	p, err := NewProduct("Product 1", 10)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, "Product 1", p.Name)
	assert.Equal(t, 10, p.Price)
}

func TestProduct_Validate_ErrIDRequired(t *testing.T) {
	product := &Product{
		ID:    entity.ID{},
		Name:  "Test Product",
		Price: 100,
	}

	err := product.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, ErrIDRequired, err)
}

func TestProduct_Validate_ErrNameRequired(t *testing.T) {
	p, err := NewProduct("", 10)
	assert.NotNil(t, err)
	assert.Nil(t, p)
	assert.Equal(t, ErrNameRequired, err)
}

func TestProduct_Validate_ErrPriceRequired(t *testing.T) {
	p, err := NewProduct("Product 1", 0)
	assert.NotNil(t, err)
	assert.Nil(t, p)
	assert.Equal(t, ErrPriceRequired, err)
}

func TestProduct_Validate_ErrPriceInvalid(t *testing.T) {
	p, err := NewProduct("Product 1", -10)
	assert.NotNil(t, err)
	assert.Nil(t, p)
	assert.Equal(t, ErrPriceInvalid, err)
}

func TestProduct_Validate(t *testing.T) {
	p, err := NewProduct("Product 1", 10)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}
