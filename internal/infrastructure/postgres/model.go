package postgres

import (
	"database/sql"
	"usermic/internal/repository"
)

type PostgresUserModel struct {
	FirstName   string         `db:"first_name"`
	MiddleName  sql.NullString `db:"middle_name"`
	LastName    string         `db:"last_name"`
	Password    string         `db:"pass_hash"`
	PhoneNumber string         `db:"phone_number"`
	Email       string         `db:"email"`
	Age         uint8          `db:"age"`
}

func recordToUserModel(ur *repository.UserRecord) *PostgresUserModel {
	middleName := sql.NullString{
		String: ur.MiddleName,
		Valid: ur.MiddleName != "",
	}

	return &PostgresUserModel{
		FirstName: ur.FirstName,
		MiddleName: middleName,
		LastName: ur.LastName,
		Password: ur.PasswordHash,
		PhoneNumber: ur.PhoneNumber,
		Email: ur.Email,
		Age: ur.Age,
	}
}

func userModelToRecord(um *PostgresUserModel) *repository.UserRecord {
	return &repository.UserRecord{
		FirstName: um.FirstName,
		MiddleName: um.MiddleName.String,
		LastName: um.LastName,
		PasswordHash: um.Password,
		PhoneNumber: um.PhoneNumber,
		Email: um.Email,
		Age: um.Age,
	}
}
