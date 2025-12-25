package handlers

import (
	"context"
	"log/slog"
	"time"
	userv1 "usermic/internal/transport/grpc/pb"
	"usermic/internal/usecase/registration"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	userv1.UnimplementedUserServiceServer
	regUc *registration.RegistrationUC
	log   *slog.Logger
	timeout time.Duration
}

func NewUserHandler(log *slog.Logger, regUc *registration.RegistrationUC, timeout time.Duration) *UserHandler {
	return &UserHandler{
		log:   log,
		regUc: regUc,
		timeout: timeout,
	}
}

func (u *UserHandler) Registration(ctx context.Context, in *userv1.RegistrationRequest) (*userv1.RegistrationResponse, error) {
	const op = "handlers.Registration"

	u.log.Info("new registration request", slog.String("op", op))

	ctx, cancel := context.WithTimeout(ctx,  u.timeout)
	defer cancel()

	regInput, err := registration.NewRegInput(
		in.GetFirstName(),
		in.GetMiddleName(),
		in.GetLastName(),
		in.GetPassword(),
		in.GetPhoneNumber(),
		in.GetEmail(),
		uint8(in.GetAge()),
	)

	if err != nil {
		u.log.Warn("invalid arguments", slog.String("op", op), slog.String("error", err.Error()))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	id, err := u.regUc.Registration(ctx, regInput)

	if err != nil {
		u.log.Warn("registration error", slog.String("op", op), slog.String("error", err.Error()))
		return nil, status.Error(codes.Canceled, err.Error())
	}

	return &userv1.RegistrationResponse{
		UserId: int32(id),
	}, nil
}