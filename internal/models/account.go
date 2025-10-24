package models

import (
	"equi_genea_account_service/internal/db"
	accountpb "equi_genea_account_service/internal/pb/api/account"
	"errors"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (a *Account) ToAccountPB() *accountpb.Account {
	return &accountpb.Account{
		Id:             a.ID.String(),
		Email:          a.Email,
		Password:       a.Password,
		CreatedAt:      timestamppb.New(a.CreatedAt),
		UpdatedAt:      timestamppb.New(a.UpdatedAt),
		LastActivityAt: timestamppb.New(a.LastActivityAt),
	}
}
