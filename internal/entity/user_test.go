package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	name := "John Doe"
	email := "j@j.com"
	pass := "123456"
	user, err := NewUser(name, email, pass)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, name, user.Name)
	assert.Equal(t, email, user.Email)
}

func TestUser_ValidetePassword(t *testing.T) {
	name := "John Doe"
	email := "j@j.com"
	pass := "123456"
	user, err := NewUser(name, email, pass)
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword(pass))
	assert.False(t, user.ValidatePassword("1234567"))
	assert.NotEqual(t, pass, user.Password)
}
