package service

import (
	"context"
	"github.com/Open-Streaming-Solutions/user-service/internal/logging"
	"github.com/Open-Streaming-Solutions/user-service/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewUserService)

type IUserService interface {
	GetUser(ctx context.Context, username string) (pgtype.UUID, error)
	CreateUser(ctx context.Context, id pgtype.UUID, username, email string) error
}

type UserService struct {
	logger logging.Logger
	db     *repository.Queries
}

func NewUserService(logger logging.Logger, db *repository.Queries) IUserService {
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

func (s *UserService) CreateUser(ctx context.Context, id pgtype.UUID, username, email string) error {
	_, err := s.db.CreateUser(ctx, repository.CreateUserParams{
		ID:       id,
		Nickname: username,
		Email:    email,
	})
	if err != nil {
		s.logger.Error("Error creating user: ", err)
		return err
	}
	return nil
}
