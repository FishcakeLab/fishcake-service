package polygon_client

import (
	"context"
	"fmt"
	"math/big"
	"net"
	"net/url"
	"time"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/FishcakeLab/fishcake-service/synchronizer/retry"
)

const (
	defaultDialTimeout    = 5 * time.Second
	defaultDialAttempts   = 5
	defaultRequestTimeout = 10 * time.Second
)

type PolygonClient interface {
	BlockHeaderByNumber(*big.Int) (*types.Header, error)
	BlockByNumber(*big.Int) (*types.Block, error)
	LatestSafeBlockHeader() (*types.Header, error)
	LatestFinalizedBlockHeader() (*types.Header, error)
	BlockHeaderByHash(common.Hash) (*types.Header, error)
	BlockByHash(common.Hash) (*types.Block, error)
	BlockHeadersByRange(*big.Int, *big.Int) ([]types.Header, error)

	TxByHash(common.Hash) (*types.Transaction, error)
	TxReceiptByHash(common.Hash) (*types.Receipt, error)

	StorageHash(common.Address, *big.Int) (common.Hash, error)
	FilterLogs(ethereum.FilterQuery) (Logs, error)

	GetBalanceByBlockNumber(address string, blockNumber *big.Int) (*big.Int, error)
	GetERC20Balance(contractAddress common.Address, ownerAddress common.Address, blocknumber *big.Int) (*big.Int, error)
	GetERC20TotalSupply(contract string, blocknumber *big.Int) (*big.Int, error)

	Close()
}

type clnt struct {
	rpc RPC
}

func DialPolygonClient(ctx context.Context, rpcUrl string) (PolygonClient, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultDialTimeout)
	defer cancel()
	bOff := retry.Exponential()
	rpcClient, err := retry.Do(ctx, defaultDialAttempts, bOff, func() (*rpc.Client, error) {
		if IsURLAvailable(rpcUrl) {
			return nil, fmt.Errorf("address unavailable (%s)", rpcUrl)
		}
		client, err := rpc.DialContext(ctx, rpcUrl)
		if err != nil {
			return nil, fmt.Errorf("failed to dial address (%s): %w", rpcUrl, err)
		}
		return client, nil
	})
	if err != nil {
		return nil, err
	}
	return &clnt{rpc: NewRPC(rpcClient)}, nil
}

func (c *clnt) BlockHeaderByHash(hash common.Hash) (*types.Header, error) {
	ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()
	var header *types.Header
	err := c.rpc.CallContext(ctxwt, &header, "eth_getBlockByHash", hash, false)
	if err != nil {
		return nil, err
	} else if header == nil {
		return nil, ethereum.NotFound
	}
	if header.Hash() != hash {
		return nil, errors.New("header mismatch")
	}
	return header, nil
}

func (c *clnt) BlockByHash(hash common.Hash) (*types.Block, error) {
	ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()
	var block *types.Block
	err := c.rpc.CallContext(ctxwt, &block, "eth_getBlockByHash", hash, false)
	if err != nil {
		return nil, err
	} else if block == nil {
		return nil, ethereum.NotFound
	}
	if block.Hash() != hash {
		return nil, errors.New("block mismatch")
	}
	return block, nil
}

func (c *clnt) LatestSafeBlockHeader() (*types.Header, error) {
	ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()
	var header *types.Header
	err := c.rpc.CallContext(ctxwt, &header, "eth_getBlockByNumber", "safe", false)
	if err != nil {
		return nil, err
	} else if header == nil {
		return nil, ethereum.NotFound
	}
	return header, nil
}

func (c *clnt) LatestFinalizedBlockHeader() (*types.Header, error) {
	ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()
	var header *types.Header
	err := c.rpc.CallContext(ctxwt, &header, "eth_getBlockByNumber", "finalized", false)
	if err != nil {
		return nil, err
	} else if header == nil {
		return nil, ethereum.NotFound
	}
	return header, nil
}

func (c *clnt) BlockHeaderByNumber(number *big.Int) (*types.Header, error) {
	ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()

	var header *types.Header
	err := c.rpc.CallContext(ctxwt, &header, "eth_getBlockByNumber", toBlockNumArg(number), false)
	if err != nil {
		return nil, err
	} else if header == nil {
		return nil, ethereum.NotFound
	}

	return header, nil
}

func (c *clnt) BlockByNumber(number *big.Int) (*types.Block, error) {
	ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()
	var block *types.Block
	err := c.rpc.CallContext(ctxwt, &block, "eth_getBlockByNumber", toBlockNumArg(number), false)
	if err != nil {
		return nil, err
	} else if block == nil {
		return nil, ethereum.NotFound
	}
	return block, nil
}

func (c *clnt) BlockHeadersByRange(startHeight, endHeight *big.Int) ([]types.Header, error) {
	if startHeight.Cmp(endHeight) == 0 {
		header, err := c.BlockHeaderByNumber(startHeight)
		if err != nil {
			return nil, err
		}
		return []types.Header{*header}, nil
	}

	count := new(big.Int).Sub(endHeight, startHeight).Uint64() + 1
	headers := make([]types.Header, count)
	batchElems := make([]rpc.BatchElem, count)

	for i := uint64(0); i < count; i++ {
		height := new(big.Int).Add(startHeight, new(big.Int).SetUint64(i))
		batchElems[i] = rpc.BatchElem{Method: "eth_getBlockByNumber", Args: []interface{}{toBlockNumArg(height), false}, Result: &headers[i]}
	}

	ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()
	err := c.rpc.BatchCallContext(ctxwt, batchElems)
	if err != nil {
		return nil, err
	}
	size := 0
	for i, batchElem := range batchElems {
		if batchElem.Error != nil {
			if size == 0 {
				return nil, batchElem.Error
			} else {
				break
			}
		} else if batchElem.Result == nil {
			break
		}
		if i > 0 && headers[i].ParentHash != headers[i-1].Hash() {
			return nil, fmt.Errorf("queried header %s does not follow parent %s", headers[i].Hash(), headers[i-1].Hash())
		}
		size = size + 1
	}
	headers = headers[:size]
	return headers, nil
}

func (c *clnt) TxByHash(hash common.Hash) (*types.Transaction, error) {
	ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()

	var tx *types.Transaction
	err := c.rpc.CallContext(ctxwt, &tx, "eth_getTransactionByHash", hash)
	if err != nil {
		return nil, err
	} else if tx == nil {
		return nil, ethereum.NotFound
	}

	return tx, nil
}

func (c *clnt) TxReceiptByHash(hash common.Hash) (*types.Receipt, error) {
	ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()
	var txReceipt *types.Receipt
	err := c.rpc.CallContext(ctxwt, &txReceipt, "eth_getTransactionReceipt", hash)
	if err != nil {
		return nil, err
	} else if txReceipt == nil {
		return nil, ethereum.NotFound
	}
	return txReceipt, nil
}

func (c *clnt) StorageHash(address common.Address, blockNumber *big.Int) (common.Hash, error) {
	ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()
	proof := struct{ StorageHash common.Hash }{}
	err := c.rpc.CallContext(ctxwt, &proof, "eth_getProof", address, nil, toBlockNumArg(blockNumber))
	if err != nil {
		return common.Hash{}, err
	}
	return proof.StorageHash, nil
}

func (c *clnt) Close() {
	c.rpc.Close()
}

type Logs struct {
	Logs          []types.Log
	ToBlockHeader *types.Header
}

func (c *clnt) FilterLogs(query ethereum.FilterQuery) (Logs, error) {
	arg, err := toFilterArg(query)
	if err != nil {
		return Logs{}, err
	}

	var logs []types.Log
	var header types.Header

	batchElems := make([]rpc.BatchElem, 2)
	batchElems[0] = rpc.BatchElem{Method: "eth_getBlockByNumber", Args: []interface{}{toBlockNumArg(query.ToBlock), false}, Result: &header}
	batchElems[1] = rpc.BatchElem{Method: "eth_getLogs", Args: []interface{}{arg}, Result: &logs}

	ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()
	err = c.rpc.BatchCallContext(ctxwt, batchElems)
	if err != nil {
		return Logs{}, err
	}

	if batchElems[0].Error != nil {
		return Logs{}, fmt.Errorf("unable to query for the `FilterQuery#ToBlock` header: %w", batchElems[0].Error)
	}

	if batchElems[1].Error != nil {
		return Logs{}, fmt.Errorf("unable to query logs: %w", batchElems[1].Error)
	}

	return Logs{Logs: logs, ToBlockHeader: &header}, nil
}

func (c *clnt) GetBalanceByBlockNumber(address string, blockNumber *big.Int) (*big.Int, error) {
	var balance *big.Int
	var err error
	ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()

	if balance, err = c.BalanceAt(ctxwt, common.HexToAddress(address), blockNumber); err != nil {
		return nil, err
	}
	return balance, nil
}

func (c *clnt) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	var result hexutil.Big
	err := c.rpc.CallContext(ctx, &result, "eth_getBalance", account, toBlockNumArg(blockNumber))

	return (*big.Int)(&result), err
}

func (c *clnt) GetERC20Balance(contractAddress common.Address, ownerAddress common.Address, blocknumber *big.Int) (*big.Int, error) {
	data := append(crypto.Keccak256([]byte("balanceOf(address)"))[:4], common.LeftPadBytes(ownerAddress.Bytes(), 32)...)
	callMsg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}
	ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()

	res, err := c.CallContract(ctxwt, callMsg, blocknumber)
	if err != nil && err.Error() != "execution reverted" {
		return nil, err
	}
	balance := new(big.Int).SetBytes(res)
	return balance, nil
}

func (c *clnt) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	var hex hexutil.Bytes
	err := c.rpc.CallContext(ctx, &hex, "eth_call", toCallArg(msg), toBlockNumArg(blockNumber))
	if err != nil {
		return nil, err
	}
	return hex, nil
}

func (c *clnt) GetERC20TotalSupply(contract string, blocknumber *big.Int) (*big.Int, error) {
	method := crypto.Keccak256([]byte("totalSupply()"))[:4]
	addr := common.HexToAddress(contract)
	callMsg := ethereum.CallMsg{
		To:   &addr,
		Data: method,
	}
	ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()

	res, err := c.CallContract(ctxwt, callMsg, blocknumber)
	if err != nil && err.Error() != "execution reverted" {
		return nil, err
	}
	supply := new(big.Int)
	supply.SetBytes(res)
	return supply, nil
}

type RPC interface {
	Close()
	CallContext(ctx context.Context, result any, method string, args ...any) error
	BatchCallContext(ctx context.Context, b []rpc.BatchElem) error
}

type rpcClient struct {
	rpc *rpc.Client
}

func NewRPC(client *rpc.Client) RPC {
	return &rpcClient{client}
}

func (c *rpcClient) Close() {
	c.rpc.Close()
}

func (c *rpcClient) CallContext(ctx context.Context, result any, method string, args ...any) error {
	err := c.rpc.CallContext(ctx, result, method, args...)
	return err
}

func (c *rpcClient) BatchCallContext(ctx context.Context, b []rpc.BatchElem) error {
	err := c.rpc.BatchCallContext(ctx, b)
	return err
}

func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	if number.Sign() >= 0 {
		return hexutil.EncodeBig(number)
	}
	return rpc.BlockNumber(number.Int64()).String()
}

func toFilterArg(q ethereum.FilterQuery) (interface{}, error) {
	arg := map[string]interface{}{"address": q.Addresses, "topics": q.Topics}
	if q.BlockHash != nil {
		arg["blockHash"] = *q.BlockHash
		if q.FromBlock != nil || q.ToBlock != nil {
			return nil, errors.New("cannot specify both BlockHash and FromBlock/ToBlock")
		}
	} else {
		if q.FromBlock == nil {
			arg["fromBlock"] = "0x0"
		} else {
			arg["fromBlock"] = toBlockNumArg(q.FromBlock)
		}
		arg["toBlock"] = toBlockNumArg(q.ToBlock)
	}
	return arg, nil
}

func toCallArg(msg ethereum.CallMsg) interface{} {
	arg := map[string]interface{}{
		"from": msg.From,
		"to":   msg.To,
	}
	if len(msg.Data) > 0 {
		arg["input"] = hexutil.Bytes(msg.Data)
	}
	if msg.Value != nil {
		arg["value"] = (*hexutil.Big)(msg.Value)
	}
	if msg.Gas != 0 {
		arg["gas"] = hexutil.Uint64(msg.Gas)
	}
	if msg.GasPrice != nil {
		arg["gasPrice"] = (*hexutil.Big)(msg.GasPrice)
	}
	return arg
}

func IsURLAvailable(address string) bool {
	u, err := url.Parse(address)
	if err != nil {
		return false
	}
	addr := u.Host
	if u.Port() == "" {
		switch u.Scheme {
		case "http", "ws":
			addr += ":80"
		case "https", "wss":
			addr += ":443"
		default:
			return true
		}
	}
	conn, err := net.DialTimeout("tcp", addr, 5*time.Second)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
