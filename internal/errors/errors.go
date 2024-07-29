package errors

import (
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

// DefaultErrRowScanPrefix is prefix for row scan errors, using in this project for wrapping QueryRow func
const DefaultErrRowScanPrefix = "number of field descriptions must equal number of"

var (
	// Validation Error
	ErrInvalidUUID = errors.New("invalid UUID: ")

	// Database Errors
	ErrUserNotFound              = errors.New("user not found")
	ErrUserAlreadyExists         = errors.New("user already exists")
	ErrSubscriptionAlreadyExists = errors.New("subscription already exists")
	ErrForeignKeyViolation       = errors.New("foreign key violation")
	ErrNotNullViolation          = errors.New("not null violation")
	ErrCheckViolation            = errors.New("check violation")
	ErrUndefinedTable            = errors.New("undefined table")
	ErrUndefinedColumn           = errors.New("undefined column")
	ErrInvalidTextRepresentation = errors.New("invalid text representation")
	ErrSyntaxError               = errors.New("syntax error")

	// SQL Error
	ErrRowScanNotMatch = errors.New("number of field descriptions do not match amount of parameters")
)

func convertPgErrorToGRPC(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			if pgErr.ConstraintName == "subscriptions_unique_constraint" {
				return status.Errorf(codes.AlreadyExists, "subscription already exists: %v", pgErr)
			}
			return status.Errorf(codes.AlreadyExists, "user already exists: %v", pgErr)
		case "23503":
			return status.Errorf(codes.InvalidArgument, "foreign key violation: %v", pgErr)
		case "23502":
			return status.Errorf(codes.InvalidArgument, "not null violation: %v", pgErr)
		case "23514":
			return status.Errorf(codes.InvalidArgument, "check violation: %v", pgErr)
		case "42P01":
			return status.Errorf(codes.InvalidArgument, "undefined table: %v", pgErr)
		case "42703":
			return status.Errorf(codes.InvalidArgument, "undefined column: %v", pgErr)
		case "22P02":
			return status.Errorf(codes.InvalidArgument, "invalid text representation: %v", pgErr)
		case "42601":
			return status.Errorf(codes.InvalidArgument, "syntax error: %v", pgErr)
		default:
			return status.Errorf(codes.Internal, "database error: %v", pgErr)
		}
	}
	return status.Errorf(codes.Internal, "unknown error: %v", err)
}

func ToGrpcError(err error) error {
	switch {
	case errors.Is(err, ErrUserNotFound):
		return status.Errorf(codes.NotFound, err.Error())
	case errors.Is(err, ErrUserAlreadyExists):
		return status.Errorf(codes.InvalidArgument, err.Error())
	case errors.Is(err, ErrUserAlreadyExists):
		return status.Errorf(codes.AlreadyExists, err.Error())
	case errors.Is(err, ErrSubscriptionAlreadyExists):
		return status.Errorf(codes.AlreadyExists, err.Error())
	case errors.Is(err, ErrForeignKeyViolation):
		return status.Errorf(codes.InvalidArgument, err.Error())
	case errors.Is(err, ErrNotNullViolation):
		return status.Errorf(codes.InvalidArgument, err.Error())
	case errors.Is(err, ErrCheckViolation):
		return status.Errorf(codes.InvalidArgument, err.Error())
	case errors.Is(err, ErrUndefinedTable):
		return status.Errorf(codes.InvalidArgument, err.Error())
	case errors.Is(err, ErrUndefinedColumn):
		return status.Errorf(codes.InvalidArgument, err.Error())
	case errors.Is(err, ErrRowScanNotMatch):
		return status.Errorf(codes.Internal, "problem with row scan, please inform the developer")
	case errors.Is(err, ErrInvalidTextRepresentation):
		return status.Errorf(codes.InvalidArgument, err.Error())
	case errors.Is(err, ErrSyntaxError):
		return status.Errorf(codes.InvalidArgument, err.Error())
	default:
		return status.Errorf(codes.Internal, "unknown error: %v", err)
	}
}

func ConvertPgError(err error) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrUserNotFound
	}
	if strings.HasPrefix(err.Error(), DefaultErrRowScanPrefix) {
		return ErrRowScanNotMatch
	}
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			if pgErr.ConstraintName == "subscriptions_unique_constraint" {
				return ErrSubscriptionAlreadyExists
			}
			return ErrUserAlreadyExists
		case "23503":
			return ErrForeignKeyViolation
		case "23502":
			return ErrNotNullViolation
		case "23514":
			return ErrCheckViolation
		case "42P01":
			return ErrUndefinedTable
		case "42703":
			return ErrUndefinedColumn
		case "22P02":
			return ErrInvalidTextRepresentation
		case "42601":
			return ErrSyntaxError
		default:
			return errors.New("database error: " + pgErr.Message)
		}
	}
	return errors.New("unknown error: " + err.Error())
}
