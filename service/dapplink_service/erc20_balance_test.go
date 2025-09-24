package dapplink_service

import (
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
)

func TestGetErc20BalanceByAddress(t *testing.T) {
	cli, err := ethclient.Dial("https://polygon-mainnet.g.alchemy.com/v2/afSCtxPWD3NE5vSjJm2GQ")
	if err != nil {
		t.Fatal(err)
	}

	ds := &DappLinkService{ethCli: cli}

	contractAddress := "0x84eBc138F4Ab844A3050a6059763D269dC9951c6" // fcc
	userAddress := "0xDdE52e4dE780A119854023817D503287aA8E8Ff0"

	balance, err := ds.GetErc20BalanceByAddress(contractAddress, userAddress)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("FCC balance: %s", balance)
	if balance == "0" {
		t.Error("expected non-zero balance")
	}
}
