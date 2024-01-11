package rpc_service

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/FishcakeLab/fishcake-service/common/enum"
	"github.com/FishcakeLab/fishcake-service/common/errors_h"
	"github.com/FishcakeLab/fishcake-service/common/global_const"
	"github.com/FishcakeLab/fishcake-service/rpc/wallet"
)

type RpcService interface {
	GetBalance(address string) *wallet.BalanceResponse
}

type rpcService struct {
	walletSercice wallet.WalletServiceClient
}

func NewRpcService(rpcUrl string) RpcService {
	conn, err := grpc.Dial(rpcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	walletServiceClient := wallet.NewWalletServiceClient(conn)
	return &rpcService{walletServiceClient}
}

func (r *rpcService) GetBalance(address string) *wallet.BalanceResponse {
	ctx := context.Background()
	balanceRequest := &wallet.BalanceRequest{Chain: global_const.Polygon, Address: address}
	balance, err := r.walletSercice.GetBalance(ctx, balanceRequest)
	if err != nil {
		errors_h.NewErrorByEnum(enum.GrpcErr)
	}
	return balance
}
