package tax

import (
	"errors"
	tax "tax/basic_testing"
)

func CalculateTaxWithError(amount float64) (float64, error) {
	if amount <= 0 {
		return 0.0, errors.New("amount must be greater than 0")
	}
	if amount >= 20000 {
		return 20.0, nil
	}
	if amount >= 1000 {
		return 10.0, nil
	}
	return 5.0, nil
}

type Repository interface {
	SaveTax(amount float64) error
}

func CalculateTaxAndSave(amount float64, repository Repository) error {
	tax := tax.CalculateTax(amount)
	return repository.SaveTax(tax)
}
