package db

import (
	"api-go-full/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("Marco", "email@gmail.com", "123")
	userDB := NewUser(db)

	err = userDB.CreateUser(user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)

	var userFound entity.User
	err = db.First(&userFound, user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.NotNil(t, userFound.Password)

}

func TestFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("Marco", "email@gmail.com", "123")
	userDB := NewUser(db)

	err = userDB.CreateUser(user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)

	userFound, err := userDB.FindByEmail("email@gmail.com")
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.NotNil(t, userFound.Password) //PAREI NA AULA 13

}
