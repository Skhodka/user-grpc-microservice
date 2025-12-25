package registration

import (
	domain "usermic/internal/domain/user"
	"usermic/internal/repository"
)

func domainToRecord(ud *domain.UserDomain) *repository.UserRecord {
	return &repository.UserRecord{
		FirstName: ud.FirstName,
		MiddleName: ud.MiddleName,
		LastName: ud.LastName,
		PasswordHash: ud.Password.GetPass(),
		PhoneNumber: ud.PhoneNumber,
		Email: ud.Email,
		Age: ud.Age,
	}
}

func recordToDomain(ur *repository.UserRecord) (*domain.UserDomain, error) {
	hashPassword, err := domain.NewHashPassword(ur.PasswordHash)

	if err != nil {
		return nil, err
	}

	return &domain.UserDomain{
		FirstName: ur.FirstName,
		MiddleName: ur.MiddleName,
		LastName: ur.LastName,
		Password: *hashPassword,
		PhoneNumber: ur.PhoneNumber,
		Email: ur.Email,
		Age: ur.Age,
	}, nil
}