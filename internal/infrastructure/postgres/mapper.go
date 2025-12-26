package postgres

import (
	"database/sql"
	domain "usermic/internal/domain/user"
)

func domainToUserModel(ud *domain.UserDomain) *PostgresUserModel {
	middleName := sql.NullString{
		String: ud.MiddleName,
		Valid:  ud.MiddleName != "",
	}

	return &PostgresUserModel{
		FirstName:   ud.FirstName,
		MiddleName:  middleName,
		LastName:    ud.LastName,
		Password:    ud.Password.GetPass(),
		PhoneNumber: ud.PhoneNumber,
		Email:       ud.Email,
		Age:         ud.Age,
	}
}

func userModelToDomain(um *PostgresUserModel) (*domain.UserDomain, error) {
	hashPass, err := domain.NewHashPassword(um.Password)

	if err != nil {
		return nil, err
	}

	return &domain.UserDomain{
		FirstName:   um.FirstName,
		MiddleName:  um.MiddleName.String,
		LastName:    um.LastName,
		Password:    *hashPass,
		PhoneNumber: um.PhoneNumber,
		Email:       um.Email,
		Age:         um.Age,
	}, nil
}
