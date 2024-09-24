package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Marco", "email@gmail.com", "123")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Marco", user.Name)
	assert.Equal(t, "email@gmail.com", user.Email)

}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Marco", "email@gmail.com", "123")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123"))
	assert.False(t, user.ValidatePassword("1234"))
	assert.NotEqual(t, "123", user.Password)
}
