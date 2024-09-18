package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mwives/go-expert/dependency-injection/product"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	productRepository := product.NewProductRepository(db)
	productUseCase := product.NewProductUseCase(productRepository)

	product, err := productUseCase.GetProduct(1)
	if err != nil {
		panic(err)
	}

	println(product.Name)
}
