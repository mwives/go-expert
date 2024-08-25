package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "j@j.com", "qwerty")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "j@j.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe", "j@j.com", "qwerty")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("qwerty"))
	assert.False(t, user.ValidatePassword("asdfgh"))
	assert.NotEqual(t, "qwerty", user.Password)
}
