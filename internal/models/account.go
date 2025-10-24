package models

import (
	"equi_genea_account_service/internal/db"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID             uuid.UUID
	Email          string
	Password       string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	LastActivityAt time.Time
}

func (a *Account) LoadFromDB(from *db.Account) error {
	if !from.ID.Valid {
		return errors.New("invalid UUID: field 'id' is null or not set")
	}

	a.ID = from.ID.Bytes
	a.Email = from.Email
	a.Password = from.Password
	a.CreatedAt = from.CreatedAt.Time
	a.UpdatedAt = from.UpdatedAt.Time
	a.LastActivityAt = from.LastActivityAt.Time

	return nil
}
