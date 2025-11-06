package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "j@email.com", "11234")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "j@email.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe", "j@email.com", "11234")

	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("11234"))
	assert.False(t, user.ValidatePassword("112345"))
	assert.NotEqual(t, "11234", user.Password)
}
