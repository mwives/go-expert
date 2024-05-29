package main

import (
	"fmt"

	"gorm.io/gorm"
)

type Product struct {
	ID           int        `gorm:"primary_key"`
	Name         string     `gorm:"type:varchar(100)"`
	Price        float64    `gorm:"type:decimal(10,2)"`
	Categories   []Category `gorm:"many2many:products_categories;"`
	SerialNumber SerialNumber
	gorm.Model
}

func CreateProduct(db *gorm.DB, product Product) {
	db.Create(&product)
}

func CreateProductsBatch(db *gorm.DB, products []Product) {
	db.Create(&products)
}

func FindProductById(db *gorm.DB, id int) Product {
	var product Product
	db.First(&product, id)
	return product
}

func FindProductByName(db *gorm.DB, name string) Product {
	var product Product
	db.Preload("Categories").First(&product, "name = ?", name)
	return product
}

func FindAllProducts(db *gorm.DB, limit int, offset int) []Product {
	var products []Product
	db.Limit(limit).Offset(offset).Preload("Category").Preload("SerialNumber").Find(&products)
	return products
}

func UpdateProduct(db *gorm.DB, product *Product) {
	db.Save(&product)
}

func DeleteProduct(db *gorm.DB, productId int) {
	var product Product
	db.First(&product, productId)
	db.Delete(&product)
}

// Helpers
func PrintProduct(product Product) {
	fmt.Printf("Product: %v; Price: %v; SerialNumber: %v\n", product.Name, product.Price, product.SerialNumber.Serial)
	println("Categories: ")
	for _, category := range product.Categories {
		fmt.Printf("- %v\n", category.Name)
	}
}
