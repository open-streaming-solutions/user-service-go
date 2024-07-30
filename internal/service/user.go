package service

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/open-streaming-solutions/user-service/internal/database"
	"github.com/open-streaming-solutions/user-service/internal/logging"
	"github.com/open-streaming-solutions/user-service/internal/repository"
	"github.com/open-streaming-solutions/user-service/internal/validation"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewUserService)

type IUserService interface {
	GetUser(ctx context.Context, username string) (pgtype.UUID, error)
	CreateUser(ctx context.Context, id string, username, email string) error
}

type UserService struct {
	logger  logging.Logger
	querier repository.QuerierWithTrx
	db      *database.Database
}

func NewUserService(logger logging.Logger, querier repository.QuerierWithTrx, db *database.Database) IUserService {
	return &UserService{
		logger:  logger,
		querier: querier,
		db:      db,
	}
}

func (s *UserService) GetUser(ctx context.Context, username string) (pgtype.UUID, error) {
	id, err := s.querier.GetUser(ctx, username)
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

	tx, err := s.db.Begin(ctx)
	if err != nil {
		s.logger.Error("Failed to begin transaction", "error", err)
		return err
	}
	defer tx.Rollback(ctx)

	queries := s.querier.WithTx(tx)

	user, err := queries.CreateUser(ctx, repository.CreateUserParams{
		ID:       *uuid,
		Username: username,
		Email:    address.String(),
	})
	if err != nil {
		s.logger.Error("Error creating user: ", err)
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		s.logger.Error("Failed to commit transaction", "error", err)
		return err
	}

	s.logger.Info("Created user", "user", user)

	return nil
}
