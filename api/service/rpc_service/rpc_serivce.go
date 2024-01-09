package rpc_service

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/FishcakeLab/fishcake-service/rpc/wallet"
)

type RpcService interface {
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

func (r *rpcService) GetChainInfo() {
	//r.walletSercice.g
}
