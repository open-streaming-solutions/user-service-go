package validation

import (
	xerrors "errors"
	"github.com/Open-Streaming-Solutions/user-service/internal/errors"
	"github.com/jackc/pgx/v5/pgtype"
	"net/mail"
)

// ValidateEmail validates the email format.
func ValidateEmail(email string) (*mail.Address, error) {
	address, err := mail.ParseAddress(email)
	if err != nil {
		return address, xerrors.Join(errors.ErrInvalidEmail, err)
	}
	return address, nil
}

// ValidateUUID validates the UUID format.
func ValidateUUID(rawUUID string) (*pgtype.UUID, error) {
	uuid := &pgtype.UUID{}
	if err := uuid.Scan(rawUUID); err != nil {
		return nil, xerrors.Join(errors.ErrInvalidUUID, err)
	}
	return uuid, nil
}
