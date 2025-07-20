package dapplink_service

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"time"

	"github.com/ethereum/go-ethereum/log"

	"github.com/FishcakeLab/fishcake-service/config"
	dapplink_api "github.com/FishcakeLab/fishcake-service/dapplink-api"
)

type DappLinkService struct {
	ethCli        *ethclient.Client
	ethDataClient *dapplink_api.EthData
}

func NewDappLinkService(conf *config.Config) (*DappLinkService, error) {
	ethCli, err := ethclient.DialContext(context.Background(), conf.PolygonRpc)
	if err != nil {
		log.Error("new eth client fail", "err", err)
		return nil, err
	}

	log.Info("init eth client success", "RpcUrl", conf.RpcUrl)

	ethDataClient, err := dapplink_api.NewEthDataClient(conf.DataApiUrl, conf.DataApiKey, time.Second*15)
	if err != nil {
		return nil, err
	}
	return &DappLinkService{
		ethCli:        ethCli,
		ethDataClient: ethDataClient,
	}, nil
}

func (ds *DappLinkService) GetErc20BalanceByAddress(contractAddress string, address string) (string, error) {
	balanceResult, err := ds.ethDataClient.GetBalanceByAddress(contractAddress, address)
	if err != nil {
		log.Error("get balance fail", "err", err)
		return "0", err
	}
	log.Info("balance result", "balance=", balanceResult, "balanceStr=", balanceResult.BalanceStr)
	balanceStr := "0"
	if balanceResult.Balance != nil && balanceResult.Balance.Int() != nil {
		balanceStr = balanceResult.Balance.Int().String()
	}
	return balanceStr, nil
}

func (ds *DappLinkService) GetPolBalanceByAddress(address string) (string, error) {
	balanceResult, err := ds.ethCli.BalanceAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		log.Error("get balance fail", "err", err)
		return "0", err
	}
	log.Info("balance result", "balanceResult", balanceResult.String())
	if balanceResult != nil {
		return balanceResult.String(), nil
	} else {
		return "0", err
	}
}
