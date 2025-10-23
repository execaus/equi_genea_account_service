package service

import "equi_genea_account_service/internal/db"

type AccountService struct {
	queries *db.Queries
}

func NewAccountService(queries *db.Queries) *AccountService {
	return &AccountService{queries: queries}
}
