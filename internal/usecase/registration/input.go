package registration

import "errors"

var (
	ErrEmptyFirstName   = errors.New("empty first name")
	ErrEmptyLastName    = errors.New("empty last name")
	ErrEmptyPassword    = errors.New("empty password")
	ErrEmptyPhoneNumber = errors.New("empty phone number")
	ErrEmptyEmail       = errors.New("empty email")
	ErrEmptyAge         = errors.New("empty age")
)

type RegInput struct {
	FirstName   string
	MiddleName  string
	LastName    string
	Password    string
	PhoneNumber string
	Email       string
	Age         uint8
}

func NewRegInput(firstName, middleName, lastName, password, phoneNumber, email string,
	age uint8,
) (*RegInput, error) {
	if !validateStringNotEmpty(firstName) {
		return nil, ErrEmptyFirstName
	}
	if !validateStringNotEmpty(lastName) {
		return nil, ErrEmptyLastName
	}
	if !validateStringNotEmpty(password) {
		return nil, ErrEmptyPassword
	}
	if !validateStringNotEmpty(phoneNumber) {
		return nil, ErrEmptyPhoneNumber
	}
	if !validateStringNotEmpty(email) {
		return nil, ErrEmptyEmail
	}
	if !validateIntNotEmpty(int(age)) {
		return nil, ErrEmptyAge
	}
	return &RegInput{
		FirstName:   firstName,
		MiddleName:  middleName,
		LastName:    lastName,
		Password:    password,
		PhoneNumber: phoneNumber,
		Email:       email,
		Age:         age,
	}, nil
}

func validateStringNotEmpty(str string) bool {
	if str == "" {
		return false
	}
	return true
}

func validateIntNotEmpty(num int) bool {
	if num == 0 {
		return false
	}
	return true
}
