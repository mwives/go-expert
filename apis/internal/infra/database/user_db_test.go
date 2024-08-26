package database

import (
	"testing"

	"github.com/mwives/go-expert/apis/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupUserDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	err = db.AutoMigrate(&entity.User{})
	assert.NoError(t, err)
	return db
}

func TestUser_Create(t *testing.T) {
	t.Parallel()
	db := setupUserDB(t)
	user, err := entity.NewUser("John Doe", "j@j.com", "qwerty")
	assert.NoError(t, err)

	userDB := NewUser(db)
	err = userDB.Create(user)
	assert.NoError(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}

func TestUser_FindByEmail(t *testing.T) {
	t.Parallel()
	db := setupUserDB(t)
	user, err := entity.NewUser("John Doe", "j@j.com", "qwerty")
	assert.NoError(t, err)

	userDB := NewUser(db)
	err = userDB.Create(user)
	assert.NoError(t, err)

	userFound, err := userDB.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}
