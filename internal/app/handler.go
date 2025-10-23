package app

import (
	"context"
	accountpb "equi_genea_account_service/internal/pb"
	"equi_genea_account_service/internal/service"
)

type AccountHandler struct {
	service *service.AccountService
	accountpb.UnimplementedAccountServiceServer
}

func (h *AccountHandler) GetAccountById(ctx context.Context, in *accountpb.GetAccountByIdRequest) (*accountpb.GetAccountByIdResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewAccountHandler(service *service.AccountService) *AccountHandler {
	return &AccountHandler{service: service}
}
