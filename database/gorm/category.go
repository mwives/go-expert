package main

import "gorm.io/gorm"

type Category struct {
	ID       int       `gorm:"primary_key"`
	Name     string    `gorm:"type:varchar(100)"`
	Products []Product `gorm:"many2many:products_categories;"`
}

func CreateCategory(db *gorm.DB, category *Category) Category {
	db.Create(&category)
	return *category
}

func FindCategoryById(db *gorm.DB, id int) Category {
	var category Category
	db.Preload("Products").Preload("Products.SerialNumber").First(&category, id)
	return category
}
