package app

import (
	"context"
	accountpb "equi_genea_account_service/internal/pb/api/account"
	"equi_genea_account_service/internal/service"

	"github.com/google/uuid"
)

type AccountHandler struct {
	service *service.AccountService
	accountpb.UnimplementedAccountServiceServer
}

func (h *AccountHandler) GetAccountById(ctx context.Context, in *accountpb.GetAccountByIdRequest) (*accountpb.GetAccountByIdResponse, error) {
	uuidId, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, err
	}

	account, err := h.service.GetAccountById(ctx, uuidId)
	if err != nil {
		return nil, err
	}

	return &accountpb.GetAccountByIdResponse{
		Account: account.ToAccountPB(),
	}, nil
}

func (h *AccountHandler) CreateAccount(ctx context.Context, in *accountpb.CreateAccountRequest) (*accountpb.CreateAccountResponse, error) {
	account, err := h.service.CreateAccount(ctx, uuid.New(), in.Email, in.PasswordHash)
	if err != nil {
		return nil, err
	}

	return &accountpb.CreateAccountResponse{
		Account: account.ToAccountPB(),
	}, nil
}

func (h *AccountHandler) IsExistByEmail(ctx context.Context, in *accountpb.IsExistByEmailRequest) (*accountpb.IsExistByEmailResponse, error) {
	isExist, err := h.service.IsExistByEmail(ctx, in.Email)
	if err != nil {
		return nil, err
	}

	return &accountpb.IsExistByEmailResponse{IsExist: isExist}, nil
}

func NewAccountHandler(service *service.AccountService) *AccountHandler {
	return &AccountHandler{service: service}
}
