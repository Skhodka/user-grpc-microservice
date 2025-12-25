package registration

import (
	"context"
	"errors"
	"log/slog"
	domain "usermic/internal/domain/user"
	"usermic/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

var (
	badId = -1
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrFailedRegUser     = errors.New("failed reg user")
)

type RegistrationUC struct {
	log  *slog.Logger
	repo repository.UserRepo
}

func NewRegistrationUC(log *slog.Logger, repo repository.UserRepo) *RegistrationUC {
	return &RegistrationUC{
		log:  log,
		repo: repo,
	}
}

func (r *RegistrationUC) Registration(ctx context.Context, regInput *RegInput) (int, error) {
	const op = "usecase.Registration"

	r.log.Info("new registration", slog.String("op", op))

	hash, err := bcrypt.GenerateFromPassword([]byte(regInput.Password), bcrypt.DefaultCost)

	if err != nil {
		r.log.Warn("failed generate from password", slog.String("op", op), slog.String("error", err.Error()))
		return badId, err
	}

	hashPassword, err := domain.NewHashPassword(string(hash))

	if err != nil {
		r.log.Warn("failed hash password", slog.String("op", op), slog.String("error", err.Error()))
		return badId, err
	}

	userDomain, err := domain.NewUserDomain(
		regInput.FirstName,
		regInput.MiddleName,
		regInput.LastName,
		*hashPassword,
		regInput.PhoneNumber,
		regInput.Email,
		regInput.Age,
	)

	if err != nil {
		r.log.Warn("failed to create user domain", slog.String("op", op), slog.String("error", err.Error()))
		return badId, err
	}

	upm, err := r.repo.FindByEmail(ctx, userDomain.Email)
	if err != nil {
		return badId, err
	}
	if upm != nil {
		r.log.Info("user already exists", slog.String("op", op))
		return badId, ErrUserAlreadyExists
	}

	upm, err = r.repo.FindByPhone(ctx, userDomain.PhoneNumber)
	if err != nil {
		return badId, err
	}
	if upm != nil {
		r.log.Info("user already exists", slog.String("op", op))
		return badId, ErrUserAlreadyExists
	}

	userRecord := domainToRecord(userDomain)

	id, err := r.repo.RegUser(ctx, userRecord)

	if err != nil {
		r.log.Info("failed reg user", slog.String("op", op))
		return badId, ErrFailedRegUser
	}

	return id, nil
}