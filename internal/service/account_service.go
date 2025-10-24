package service

import (
	"context"
	"equi_genea_account_service/internal/db"
	"equi_genea_account_service/internal/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type AccountService struct {
	queries *db.Queries
}

func (s *AccountService) GetAccountById(ctx context.Context, id uuid.UUID) (*models.Account, error) {
	uuidId := pgtype.UUID{
		Bytes: id,
		Valid: true,
	}

	dbAccount, err := s.queries.GetAccountFromId(ctx, uuidId)
	if err != nil {
		return nil, err
	}

	outputModel := models.Account{}
	if err = outputModel.LoadFromDB(&dbAccount); err != nil {
		return nil, err
	}

	return &outputModel, nil
}

func NewAccountService(queries *db.Queries) *AccountService {
	return &AccountService{queries: queries}
}
