package dapplink_api

import (
	"time"

	"github.com/ethereum/go-ethereum/log"

	"github.com/dapplink-labs/chain-explorer-api/common/account"
	"github.com/dapplink-labs/chain-explorer-api/common/chain"
	"github.com/dapplink-labs/chain-explorer-api/explorer/etherscan"
)

type EthData struct {
	EthDataCli *etherscan.ChainExplorerAdaptor
}

func NewEthDataClient(baseUrl, apiKey string, timeout time.Duration) (*EthData, error) {
	etherscanCli, err := etherscan.NewChainExplorerAdaptor(apiKey, baseUrl, false, time.Duration(timeout))
	if err != nil {
		log.Error("New etherscan client fail", "err", err)
		return nil, err
	}
	return &EthData{EthDataCli: etherscanCli}, err
}

func (ed *EthData) GetTxByAddress(page, pagesize uint64, address string, action account.ActionType) (*account.TransactionResponse[account.AccountTxResponse], error) {
	request := &account.AccountTxRequest{
		PageRequest: chain.PageRequest{
			Page:  page,
			Limit: pagesize,
		},
		Action:  action,
		Address: address,
	}
	txData, err := ed.EthDataCli.GetTxByAddress(request)
	if err != nil {
		return nil, err
	}
	return txData, nil
}

func (ed *EthData) GetBalanceByAddress(contractAddr, address string) (*account.AccountBalanceResponse, error) {
	accountItem := []string{address}
	symbol := []string{"ETH"}
	contractAddress := []string{contractAddr}
	protocolType := []string{""}
	page := []string{"1"}
	limit := []string{"10"}
	acbr := &account.AccountBalanceRequest{
		ChainShortName:  "ETH",
		ExplorerName:    "etherescan",
		Account:         accountItem,
		Symbol:          symbol,
		ContractAddress: contractAddress,
		ProtocolType:    protocolType,
		Page:            page,
		Limit:           limit,
	}
	etherscanResp, err := ed.EthDataCli.GetAccountBalance(acbr)
	if err != nil {
		log.Error("get account balance error", "err", err)
		return nil, err
	}
	return etherscanResp, nil
}
