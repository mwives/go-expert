package main

import "gorm.io/gorm"

type SerialNumber struct {
	ID        int `gorm:"primary_key"`
	ProductID int
	Serial    string `gorm:"type:varchar(100)"`
}

func CreateSerialNumber(db *gorm.DB, serialNumber SerialNumber) int {
	db.Create(&serialNumber)
	return serialNumber.ID
}
