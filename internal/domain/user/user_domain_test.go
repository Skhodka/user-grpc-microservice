package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUserDomain_ValidInput(t *testing.T) {
	hashPassword, _ := NewHashPassword("any_pass")

	userDomain, err := NewUserDomain(
		"Ekaterina",
		"Middle",
		"Rabova",
		*hashPassword,
		"12345678909",
		"mail@mail.ru",
		18,
	)

	assert.NoError(t, err)
	assert.Equal(t, "Ekaterina", userDomain.FirstName)
	assert.Equal(t, "Middle", userDomain.MiddleName)
	assert.Equal(t, "Rabova", userDomain.LastName)
	assert.Equal(t, *hashPassword, userDomain.Password)
	assert.Equal(t, "12345678909", userDomain.PhoneNumber)
	assert.Equal(t, "mail@mail.ru", userDomain.Email)
	assert.Equal(t, uint8(18), userDomain.Age)
}

func TestNewUserDomain_InvalidPhone(t *testing.T) {
	hashPassword, _ := NewHashPassword("any_pass")

	_, err := NewUserDomain(
		"Ekaterina",
		"Middle",
		"Rabova",
		*hashPassword,
		"1234567890999",
		"mail@mail.ru",
		18,
	)

	assert.Error(t, err, ErrInvalidPhoneNumber)

	_, err = NewUserDomain(
		"Ekaterina",
		"Middle",
		"Rabova",
		*hashPassword,
		"1234567909",
		"mail@mail.ru",
		18,
	)

	assert.Error(t, err, ErrInvalidPhoneNumber)
}

func TestNewUserDomain_InvalidEmail(t *testing.T) {
	hashPassword, _ := NewHashPassword("any_pass")

	_, err := NewUserDomain(
		"Ekaterina",
		"Middle",
		"Rabova",
		*hashPassword,
		"12345678909",
		"mailmail.ru",
		18,
	)

	assert.Error(t, err, ErrInvalidEmail)

	_, err = NewUserDomain(
		"Ekaterina",
		"Middle",
		"Rabova",
		*hashPassword,
		"12345678909",
		"mai lmail.ru",
		18,
	)

	assert.Error(t, err, ErrInvalidEmail)
}

func TestNewUserDomain_InvalidAge(t *testing.T) {
	hashPassword, _ := NewHashPassword("any_pass")

	_, err := NewUserDomain(
		"Ekaterina",
		"Middle",
		"Rabova",
		*hashPassword,
		"12345678909",
		"mail@mail.ru",
		17,
	)

	assert.Error(t, err, ErrInvalidAge)

	_, err = NewUserDomain(
		"Ekaterina",
		"Middle",
		"Rabova",
		*hashPassword,
		"12345678909",
		"mail@mail.ru",
		200,
	)

	assert.Error(t, err, ErrInvalidAge)
}