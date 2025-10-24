package app

import (
	"context"
	accountpb "equi_genea_account_service/internal/pb/api/account"
	"equi_genea_account_service/internal/service"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
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
		Account: &accountpb.Account{
			Id:             account.ID.String(),
			Email:          account.Email,
			Password:       account.Password,
			CreatedAt:      timestamppb.New(account.CreatedAt),
			UpdatedAt:      timestamppb.New(account.UpdatedAt),
			LastActivityAt: timestamppb.New(account.LastActivityAt),
		},
	}, nil
}

func (h *AccountHandler) CreateAccount(ctx context.Context, in *accountpb.CreateAccountRequest) (*accountpb.CreateAccountResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (h *AccountHandler) IsExistByEmail(ctx context.Context, in *accountpb.IsExistByEmailRequest) (*accountpb.IsExistByEmailResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewAccountHandler(service *service.AccountService) *AccountHandler {
	return &AccountHandler{service: service}
}
