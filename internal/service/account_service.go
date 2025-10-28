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

func (s *AccountService) GetAccountByEmail(ctx context.Context, email string) (*models.Account, error) {
	dbAccount, err := s.queries.GetAccountByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	outputModel := models.Account{}
	if err = outputModel.LoadFromDB(&dbAccount); err != nil {
		return nil, err
	}

	return &outputModel, nil
}

func (s *AccountService) GetAccountById(ctx context.Context, id uuid.UUID) (*models.Account, error) {
	uuidId := pgtype.UUID{
		Bytes: id,
		Valid: true,
	}

	dbAccount, err := s.queries.GetAccountById(ctx, uuidId)
	if err != nil {
		return nil, err
	}

	outputModel := models.Account{}
	if err = outputModel.LoadFromDB(&dbAccount); err != nil {
		return nil, err
	}

	return &outputModel, nil
}

func (s *AccountService) IsExistByEmail(ctx context.Context, email string) (bool, error) {
	isExist, err := s.queries.IsExistByEmail(ctx, email)
	if err != nil {
		return false, err
	}

	return isExist, nil
}

func (s *AccountService) CreateAccount(ctx context.Context, email, password string) (*models.Account, error) {
	accountID := pgtype.UUID{
		Bytes: uuid.New(),
		Valid: true,
	}

	dbAccount, err := s.queries.CreateAccount(ctx, db.CreateAccountParams{
		ID:       accountID,
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	account := &models.Account{}
	if err = account.LoadFromDB(&dbAccount); err != nil {
		return nil, err
	}

	return account, nil
}

func NewAccountService(queries *db.Queries) *AccountService {
	return &AccountService{queries: queries}
}
