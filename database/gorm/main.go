package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// electronicsCategory := CreateCategory(db, &Category{Name: "Electronics"})
	// audioCategory := CreateCategory(db, &Category{Name: "Audio"})
	// electronicsCategory := FindCategoryById(db, 1)
	// audioCategory := FindCategoryById(db, 2)

	// CreateProduct(db, Product{
	// 	Name:       "Truthear Zero:RED",
	// 	Price:      54.0,
	// 	Categories: []Category{electronicsCategory, audioCategory},
	// })
	zeroRed := FindProductByName(db, "Truthear Zero:RED")
	PrintProduct(zeroRed)

	// CreateProductsBatch(db, []Product{
	// 	{Name: "Corsair Virtuoso", Price: 199.99, CategoryId: electronicsCategoryId},
	// 	{Name: "SteelSeries Arctis 7", Price: 179.99, CategoryId: electronicsCategoryId},
	// 	{Name: "Truthear Zero:RED", Price: 54.0, CategoryId: electronicsCategoryId},
	// })

	// zeroRed := FindProductByName(db, "Truthear Zero:RED")
	// fmt.Printf("%+v\n", zeroRed)

	// products := FindAllProducts(db, 10, 0)
	// for _, product := range products {
	// 	CreateSerialNumber(db, SerialNumber{ProductID: product.ID, Serial: "1234567890"})
	// }
	// for _, product := range products {
	// 	fmt.Printf(
	// 		"Product: %v; Price: %v; Category: %v; SerialNumber: %v\n",
	// 		product.Name, product.Price, product.Category.Name, product.SerialNumber.Serial,
	// 	)
	// }

	// category := FindCategoryById(db, 2)
	// fmt.Printf("Category: %v\n", category.Name)
	// for _, product := range category.Products {
	// 	fmt.Printf("Product: %v; Price: %v; SerialNumber: %v\n", product.Name, product.Price, product.SerialNumber.Serial)
	// }

	// product := FindProductById(db, 12)
	// product.Name = "PRO X SUPERLIGHT 2"
	// product.Price = 159.0
	// UpdateProduct(db, &product)
	// fmt.Printf("Product: %v; Price: %v\n", product.Name, product.Price)

	// DeleteProduct(db, 12)
}
