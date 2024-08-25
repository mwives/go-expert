package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test with assert
func TestCalculateTaxWithError(t *testing.T) {
	tax, err := CalculateTaxWithError(1000.0)
	assert.Nil(t, err)
	assert.Equal(t, 10.0, tax)

	tax, err = CalculateTaxWithError(0)
	assert.Error(t, err, "amount must be greater than 0")
	assert.Equal(t, 0.0, tax)
}

// Test with mock
func TestCalculateTaxAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil)
	repository.On("SaveTax", 0.0).Return(errors.New("save tax error"))

	err := CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0.0, repository)
	assert.Error(t, err, "save tax error")

	repository.AssertExpectations(t)
	repository.AssertNumberOfCalls(t, "SaveTax", 2)
	repository.AssertCalled(t, "SaveTax", 10.0)
	repository.AssertCalled(t, "SaveTax", 0.0)
}
