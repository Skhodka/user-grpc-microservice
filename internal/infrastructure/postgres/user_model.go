package postgres

import (
	"database/sql"
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
