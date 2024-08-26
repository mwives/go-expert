package entity

import (
	"errors"
	"time"

	"github.com/mwives/go-expert/apis/pkg/entity"
)

var (
	ErrIDRequired    = errors.New("id is required")
	ErrIDInvalid     = errors.New("id is invalid")
	ErrNameRequired  = errors.New("name is required")
	ErrPriceRequired = errors.New("price is required")
	ErrPriceInvalid  = errors.New("price is invalid")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
	err := product.Validate()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "00000000-0000-0000-0000-000000000000" {
		return ErrIDRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrIDInvalid
	}
	if p.Name == "" {
		return ErrNameRequired
	}
	if p.Price == 0 {
		return ErrPriceRequired
	}
	if p.Price < 0 {
		return ErrPriceInvalid
	}
	return nil
}
