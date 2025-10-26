package rpc_service

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"github.com/FishcakeLab/fishcake-service/rpc/account"
	"google.golang.org/grpc/credentials/insecure"
)

type RpcService interface {
	GetSupportChains(ctx context.Context, in *account.SupportChainsRequest) (*account.SupportChainsResponse, error)
	ConvertAddress(ctx context.Context, in *account.ConvertAddressRequest) (*account.ConvertAddressResponse, error)
	ValidAddress(ctx context.Context, in *account.ValidAddressRequest) (*account.ValidAddressResponse, error)
	GetBlockByNumber(ctx context.Context, in *account.BlockNumberRequest) (*account.BlockResponse, error)
	GetBlockByHash(ctx context.Context, in *account.BlockHashRequest) (*account.BlockResponse, error)
	GetBlockHeaderByHash(ctx context.Context, in *account.BlockHeaderHashRequest) (*account.BlockHeaderResponse, error)
	GetBlockHeaderByNumber(ctx context.Context, in *account.BlockHeaderNumberRequest) (*account.BlockHeaderResponse, error)
	GetBlockHeaderByRange(ctx context.Context, in *account.BlockByRangeRequest) (*account.BlockByRangeResponse, error)
	GetAccount(ctx context.Context, in *account.AccountRequest) (*account.AccountResponse, error)
	GetFee(ctx context.Context, in *account.FeeRequest) (*account.FeeResponse, error)
	SendTx(ctx context.Context, in *account.SendTxRequest) (*account.SendTxResponse, error)
	GetTxByAddress(ctx context.Context, in *account.TxAddressRequest) (*account.TxAddressResponse, error)
	GetTxByHash(ctx context.Context, in *account.TxHashRequest) (*account.TxHashResponse, error)
	CreateUnSignTransaction(ctx context.Context, in *account.UnSignTransactionRequest) (*account.UnSignTransactionResponse, error)
	BuildSignedTransaction(ctx context.Context, in *account.SignedTransactionRequest) (*account.SignedTransactionResponse, error)
	DecodeTransaction(ctx context.Context, in *account.DecodeTransactionRequest) (*account.DecodeTransactionResponse, error)
	VerifySignedTransaction(ctx context.Context, in *account.VerifyTransactionRequest) (*account.VerifyTransactionResponse, error)
	GetExtraData(ctx context.Context, in *account.ExtraDataRequest) (*account.ExtraDataResponse, error)
}

type rpcService struct {
	accountService account.WalletAccountServiceClient
}

func NewRpcService(rpcUrl string) RpcService {
	conn, err := grpc.Dial(rpcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	accountServiceClient := account.NewWalletAccountServiceClient(conn)
	return &rpcService{accountService: accountServiceClient}
}

func (r *rpcService) GetSupportChains(ctx context.Context, in *account.SupportChainsRequest) (*account.SupportChainsResponse, error) {
	resp, err := r.accountService.GetSupportChains(ctx, in)
	if err != nil {
		return nil, nil
	}
	return resp, nil
}

func (r *rpcService) ConvertAddress(ctx context.Context, in *account.ConvertAddressRequest) (*account.ConvertAddressResponse, error) {
	resp, err := r.accountService.ConvertAddress(ctx, in)
	if err != nil {
		return nil, nil
	}
	return resp, nil
}

func (r *rpcService) ValidAddress(ctx context.Context, in *account.ValidAddressRequest) (*account.ValidAddressResponse, error) {
	resp, err := r.accountService.ValidAddress(ctx, in)
	if err != nil {
		return nil, nil
	}
	return resp, nil
}

func (r *rpcService) GetBlockByNumber(ctx context.Context, in *account.BlockNumberRequest) (*account.BlockResponse, error) {
	resp, err := r.accountService.GetBlockByNumber(ctx, in)
	if err != nil {
		return nil, nil
	}
	return resp, nil
}

func (r *rpcService) GetBlockByHash(ctx context.Context, in *account.BlockHashRequest) (*account.BlockResponse, error) {
	resp, err := r.accountService.GetBlockByHash(ctx, in)
	if err != nil {
		return nil, nil
	}
	return resp, nil
}

func (r *rpcService) GetBlockHeaderByHash(ctx context.Context, in *account.BlockHeaderHashRequest) (*account.BlockHeaderResponse, error) {
	resp, err := r.accountService.GetBlockHeaderByHash(ctx, in)
	if err != nil {
		return nil, nil
	}
	return resp, nil
}

func (r *rpcService) GetBlockHeaderByNumber(ctx context.Context, in *account.BlockHeaderNumberRequest) (*account.BlockHeaderResponse, error) {
	resp, err := r.accountService.GetBlockHeaderByNumber(ctx, in)
	if err != nil {
		return nil, nil
	}
	return resp, nil
}

func (r *rpcService) GetBlockHeaderByRange(ctx context.Context, in *account.BlockByRangeRequest) (*account.BlockByRangeResponse, error) {
	resp, err := r.accountService.GetBlockHeaderByRange(ctx, in)
	if err != nil {
		return nil, nil
	}
	return resp, nil
}

func (r *rpcService) GetAccount(ctx context.Context, in *account.AccountRequest) (*account.AccountResponse, error) {
	resp, err := r.accountService.GetAccount(ctx, in)
	if err != nil {
		return nil, nil
	}
	return resp, nil
}

func (r *rpcService) GetFee(ctx context.Context, in *account.FeeRequest) (*account.FeeResponse, error) {
	resp, err := r.accountService.GetFee(ctx, in)
	if err != nil {
		return nil, nil
	}
	return resp, nil
}

func (r *rpcService) SendTx(ctx context.Context, in *account.SendTxRequest) (*account.SendTxResponse, error) {
	resp, err := r.accountService.SendTx(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *rpcService) GetTxByAddress(ctx context.Context, in *account.TxAddressRequest) (*account.TxAddressResponse, error) {
	resp, err := r.accountService.GetTxByAddress(ctx, in)
	if err != nil {
		return nil, nil
	}
	return resp, nil
}

func (r *rpcService) GetTxByHash(ctx context.Context, in *account.TxHashRequest) (*account.TxHashResponse, error) {
	resp, err := r.accountService.GetTxByHash(ctx, in)
	if err != nil {
		return nil, nil
	}
	return resp, nil
}

func (r *rpcService) CreateUnSignTransaction(ctx context.Context, in *account.UnSignTransactionRequest) (*account.UnSignTransactionResponse, error) {
	resp, err := r.accountService.BuildUnSignTransaction(ctx, in)
	if err != nil {
		return nil, nil
	}
	return resp, nil
}

func (r *rpcService) BuildSignedTransaction(ctx context.Context, in *account.SignedTransactionRequest) (*account.SignedTransactionResponse, error) {
	resp, err := r.accountService.BuildSignedTransaction(ctx, in)
	if err != nil {
		return nil, nil
	}
	return resp, nil
}

func (r *rpcService) DecodeTransaction(ctx context.Context, in *account.DecodeTransactionRequest) (*account.DecodeTransactionResponse, error) {
	resp, err := r.accountService.DecodeTransaction(ctx, in)
	if err != nil {
		return nil, nil
	}
	return resp, nil
}

func (r *rpcService) VerifySignedTransaction(ctx context.Context, in *account.VerifyTransactionRequest) (*account.VerifyTransactionResponse, error) {
	resp, err := r.accountService.VerifySignedTransaction(ctx, in)
	if err != nil {
		return nil, nil
	}
	return resp, nil
}

func (r *rpcService) GetExtraData(ctx context.Context, in *account.ExtraDataRequest) (*account.ExtraDataResponse, error) {
	resp, err := r.accountService.GetExtraData(ctx, in)
	if err != nil {
		return nil, nil
	}
	return resp, nil
}
