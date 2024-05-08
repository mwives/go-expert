package main

import "fmt"

type Address struct {
	Street string
	Number int
	City   string
	State  string
}

type User struct {
	Name   string
	Age    int
	Active bool
	Address
}

func (u *User) Deactivate() {
	u.Active = false
	fmt.Printf("User %s was deactivated\n", u.Name)
}

// Structs
func main() {
	ivonei := User{
		Name:   "Ivonei",
		Age:    20,
		Active: true,
	}

	ivonei.Age = 21
	ivonei.Address.City = "Hollyudi"

	ivonei.Deactivate()

	fmt.Printf("User %s has Active = %t", ivonei.Name, ivonei.Active)
}
