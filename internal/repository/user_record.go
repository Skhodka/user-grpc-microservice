package repository

type UserRecord struct {
	ID           int
	FirstName    string
	MiddleName   string
	LastName     string
	PasswordHash string
	PhoneNumber  string
	Email        string
	Age          uint8
}