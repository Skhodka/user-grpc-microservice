package postgres

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	domain "usermic/internal/domain/user"
)

var (
	ErrSelect = errors.New("failed to get data")
	ErrInsert = errors.New("failed to insert data")
)

var (
	badId = -1
)

type PostgresStorage struct {
	log *slog.Logger
	db  *sql.DB
}

func NewPostgresStorage(log *slog.Logger, db *sql.DB) *PostgresStorage {
	return &PostgresStorage{
		log: log,
		db:  db,
	}
}

func (p *PostgresStorage) RegUser(ctx context.Context, ud *domain.UserDomain) (int, error) {
	const op = "postgres.RegUser"

	p.log.Info("new registration", slog.String("op", op))

	pm := domainToUserModel(ud)

	row := p.db.QueryRowContext(
		ctx,
		`INSERT INTO users (first_name, middle_name, last_name, pass_hash, phone_number, email, age) 
		VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id;`,
		pm.FirstName, pm.MiddleName, pm.LastName, pm.Password, pm.PhoneNumber, pm.Email, pm.Age,
	)

	var id int = 0

	err := row.Scan(&id)

	if err != nil {
		p.log.Warn("failed insert", slog.String("op", op), slog.String("error", err.Error()))
		return badId, ErrInsert
	}

	return id, nil
}

func (p *PostgresStorage) FindByEmail(ctx context.Context, email string) (*domain.UserDomain, error) {
	const op = "postgres.FindByEmail"

	p.log.Info("search for a user by ID", slog.String("op", op))
	var pm *PostgresUserModel = &PostgresUserModel{}

	row := p.db.QueryRowContext(ctx, "SELECT first_name, middle_name, last_name, pass_hash, phone_number, email, age FROM users WHERE email = $1", email)

	err := row.Scan(
		&pm.FirstName,
		&pm.MiddleName,
		&pm.LastName,
		&pm.Password,
		&pm.PhoneNumber,
		&pm.Email,
		&pm.Age,
	)

	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			p.log.Warn("failed select", slog.String("op", op), slog.String("error", err.Error()))
			return nil, ErrSelect
		} else {
			return nil, nil
		}
	}

	userOut, err := userModelToDomain(pm)

	if err != nil {
		p.log.Warn("failed to create domain", slog.String("op", op))
		return nil, err
	}

	return userOut, nil
}

func (p *PostgresStorage) FindByPhone(ctx context.Context, phone string) (*domain.UserDomain, error) {
	const op = "postgres.FindByPhone"

	p.log.Info("search for a user by ID", slog.String("op", op))
	var pm *PostgresUserModel = &PostgresUserModel{}

	row := p.db.QueryRowContext(ctx, "SELECT first_name, middle_name, last_name, pass_hash, phone_number, email, age FROM users WHERE phone_number = $1", phone)

	err := row.Scan(
		&pm.FirstName,
		&pm.MiddleName,
		&pm.LastName,
		&pm.Password,
		&pm.PhoneNumber,
		&pm.Email,
		&pm.Age,
	)

	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			p.log.Warn("failed select", slog.String("op", op), slog.String("error", err.Error()))
			return nil, ErrSelect
		} else {
			return nil, nil
		}
	}
	userOut, err := userModelToDomain(pm)

	if err != nil {
		p.log.Warn("failed to create domain", slog.String("op", op))
		return nil, err
	}

	return userOut, nil
}
