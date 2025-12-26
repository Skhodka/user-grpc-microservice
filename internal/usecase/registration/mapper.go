package registration

import (
	domain "usermic/internal/domain/user"
)

func inputToDomain(ri *RegInput, hp *domain.HashPassword) (*domain.UserDomain, error) {
	return &domain.UserDomain{
		FirstName: ri.FirstName,
		MiddleName: ri.MiddleName,
		LastName: ri.LastName,
		Password: *hp,
		PhoneNumber: ri.PhoneNumber,
		Email: ri.Email,
		Age: ri.Age,
	}, nil
}