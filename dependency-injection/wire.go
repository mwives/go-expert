//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/mwives/go-expert/dependency-injection/product"
)

func NewProductUseCase(db *sql.DB) *product.ProductUseCase {
	wire.Build(
		product.NewProductRepository,
		product.NewProductUseCase,
	)
	return &product.ProductUseCase{}
}
