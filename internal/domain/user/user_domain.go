package domain

import (
	"errors"
	"net/mail"
)

var (
	ErrInvalidPhoneNumber = errors.New("invalid phone number")
	ErrInvalidEmail       = errors.New("invalid email format")
	ErrInvalidAge         = errors.New("unacceptable age")
)

type UserDomain struct {
	FirstName   string
	MiddleName  string
	LastName    string
	Password    HashPassword
	PhoneNumber string
	Email       string
	Age         uint8
}

func NewUserDomain(
	firstName, middleName, lastName string, password HashPassword, phoneNumber, email string,
	age uint8,
) (*UserDomain, error) {
	if !validatePhoneNumber(phoneNumber) {
		return nil, ErrInvalidPhoneNumber
	}

	if !validateEmailFormat(email) {
		return nil, ErrInvalidEmail
	}

	if age < 18 || age > 100 {
		return nil, ErrInvalidAge
	}

	return &UserDomain{
		FirstName:   firstName,
		MiddleName:  middleName,
		LastName:    lastName,
		Password:    password,
		PhoneNumber: phoneNumber,
		Email:       email,
		Age:         uint8(age),
	}, nil
}

func validatePhoneNumber(phone string) bool {
	phoneB := []byte(phone)

	if len(phoneB) != 11 {
		return false
	}

	for _, s := range phone {
		if s < '0' || s > '9' {
			return false
		}
	}

	return true
}

func validateEmailFormat(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}

	return true
}
