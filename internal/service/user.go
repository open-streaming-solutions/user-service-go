package service

import (
	"context"
	"github.com/Open-Streaming-Solutions/user-service/internal/logging"
	"github.com/Open-Streaming-Solutions/user-service/internal/repository"
	"github.com/Open-Streaming-Solutions/user-service/internal/validation"
	"github.com/jackc/pgx/v5/pgtype"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewUserService)

type IUserService interface {
	GetUser(ctx context.Context, username string) (pgtype.UUID, error)
	CreateUser(ctx context.Context, id string, username, email string) error
}

type UserService struct {
	logger logging.Logger
	db     repository.Querier
}

func NewUserService(logger logging.Logger, db repository.Querier) IUserService {
	return &UserService{
		logger: logger,
		db:     db,
	}
}

func (s *UserService) GetUser(ctx context.Context, username string) (pgtype.UUID, error) {
	id, err := s.db.GetUser(ctx, username)
	if err != nil {
		s.logger.Error("Error getting user from DB: ", err)
		return pgtype.UUID{}, err
	}
	return id.ID, nil
}

func (s *UserService) CreateUser(ctx context.Context, id string, username, email string) error {
	uuid, err := validation.ValidateUUID(id)
	if err != nil {
		s.logger.Error("Failed to scan UUID", "error", err)
		return err
	}

	address, err := validation.ValidateEmail(email)
	if err != nil {
		s.logger.Error("Failed to scan Email", "error", err)
		return err
	}

	_, err = s.db.CreateUser(ctx, repository.CreateUserParams{
		ID:       *uuid,
		Username: username,
		Email:    address.String(),
	})
	if err != nil {
		s.logger.Error("Error creating user: ", err)
		return err
	}

	return nil
}
