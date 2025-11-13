package node

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"net"
	"net/url"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/FishcakeLab/fishcake-service/common/global_const"
	"github.com/FishcakeLab/fishcake-service/synchronizer/retry"
)

const (
	// defaultDialTimeout is default duration the processor will wait on
	// startup to make a connection to the backend
	defaultDialTimeout = 5 * time.Second

	// defaultDialAttempts is the default attempts a connection will be made
	// before failing
	defaultDialAttempts = 5

	// defaultRequestTimeout is the default duration the processor will
	// wait for a request to be fulfilled
	defaultRequestTimeout = 100 * time.Second
)

type EthClient interface {
	BlockHeaderByNumber(*big.Int) (*types.Header, error)
	LatestSafeBlockHeader() (*types.Header, error)
	LatestFinalizedBlockHeader() (*types.Header, error)
	BlockHeaderByHash(common.Hash) (*types.Header, error)
	BlockHeadersByRange(*big.Int, *big.Int, uint) ([]types.Header, error)

	TxByHash(common.Hash) (*types.Transaction, error)

	TxReceiptByHash(common.Hash) (*types.Receipt, error)

	StorageHash(common.Address, *big.Int) (common.Hash, error)
	FilterLogs(ethereum.FilterQuery) (Logs, error)

	BlockByNumber(number *big.Int) (*types.Block, error)

	// Close closes the underlying RPC connection.
	// RPC close does not return any errors, but does shut down e.g. a websocket connection.
	Close()
}

type clnt struct {
	rpc RPC
	ec  *ethclient.Client // 新增一个 ethclient
}

func DialEthClient(ctx context.Context, rpcUrl string) (EthClient, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultDialTimeout)
	defer cancel()

	bOff := retry.Exponential()
	rpcClient, err := retry.Do(ctx, defaultDialAttempts, bOff, func() (*rpc.Client, error) {
		if !IsURLAvailable(rpcUrl) {
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

	// 基于同一个底层连接创建 ethclient
	ethCl := ethclient.NewClient(rpcClient)

	// 初始化并返回
	return &clnt{
		rpc: NewRPC(rpcClient),
		ec:  ethCl,
	}, nil
}

// BlockHeaderByHash retrieves the block header attributed to the supplied hash
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

	// sanity check on the data returned
	if header.Hash() != hash {
		return nil, errors.New("header mismatch")
	}

	return header, nil
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

// BlockHeaderByNumber retrieves the block header attributed to the supplied height
func (c *clnt) BlockHeaderByNumber(number *big.Int) (*types.Header, error) {
	ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()

	var header *types.Header
	err := c.rpc.CallContext(ctxwt, &header, "eth_getBlockByNumber", toBlockNumArg(number), false)
	if err != nil {
		log.Error("Call eth_getBlockByNumber method fail", "err", err)
		return nil, err
	} else if header == nil {
		log.Info("header not found")
		return nil, ethereum.NotFound
	}

	return header, nil
}

func (c *clnt) BlockByNumber(number *big.Int) (*types.Block, error) {
	ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()

	block, err := c.ec.BlockByNumber(ctxwt, number)
	if err != nil {
		return nil, fmt.Errorf("ethclient.BlockByNumber failed: %w", err)
	}
	if block == nil {
		return nil, fmt.Errorf("ethclient returned nil block at number=%v", number)
	}
	return block, nil
}

// BlockHeadersByRange will retrieve block headers within the specified range -- inclusive. No restrictions
// are placed on the range such as blocks in the "latest", "safe" or "finalized" states. If the specified
// range is too large, `endHeight > latest`, the resulting list is truncated to the available headers
// BlockHeadersByRange will retrieve block headers within the specified range -- inclusive.
// No restrictions are placed on the range such as blocks in the "latest", "safe" or "finalized" states.
// If the specified range is too large, `endHeight > latest`, the resulting list is truncated
// to the available headers.
func (c *clnt) BlockHeadersByRange(startHeight, endHeight *big.Int, chainId uint) ([]types.Header, error) {
	// ---- 基础参数保护，防止奇怪输入导致 panic ----
	if startHeight == nil || endHeight == nil {
		return nil, fmt.Errorf("startHeight or endHeight is nil")
	}
	if startHeight.Cmp(endHeight) > 0 {
		return nil, fmt.Errorf("startHeight (%v) > endHeight (%v)", startHeight, endHeight)
	}

	// 无范围的情况，直接单查
	if startHeight.Cmp(endHeight) == 0 {
		header, err := c.BlockHeaderByNumber(startHeight)
		if err != nil {
			return nil, err
		}
		if header == nil {
			return nil, fmt.Errorf("no header found at %v", startHeight)
		}
		return []types.Header{*header}, nil
	}

	// 计算区块个数（end - start + 1）
	diff := new(big.Int).Sub(endHeight, startHeight)
	if !diff.IsUint64() {
		return nil, fmt.Errorf("block range too large: [%v, %v]", startHeight, endHeight)
	}
	count := diff.Uint64() + 1
	if count == 0 {
		return nil, fmt.Errorf("block range size is zero")
	}

	// headers: 真正要返回的结果
	headers := make([]types.Header, count)
	// batchElems: 承载 RPC 调用的错误与原始返回
	batchElems := make([]rpc.BatchElem, count)

	// --------- 分链逻辑：非 Polygon 直接 BatchCall，Polygon 单 RPC 并发 ---------

	if chainId != uint(global_const.PolygonChainId) {
		//  非 Polygon：标准 batch 模式
		for i := uint64(0); i < count; i++ {
			height := new(big.Int).Add(startHeight, new(big.Int).SetUint64(i))
			batchElems[i] = rpc.BatchElem{
				Method: "eth_getBlockByNumber",
				Args:   []interface{}{toBlockNumArg(height), false},
				Result: &headers[i], // 结果直接填充到 headers[i]
			}
		}

		ctx, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
		defer cancel()

		if err := c.rpc.BatchCallContext(ctx, batchElems); err != nil {
			return nil, err
		}
	} else {
		//  Polygon：分组 + goroutine 并发 + 单 RPC 调用
		const groupSize = 100

		var wg sync.WaitGroup
		for start := 0; start < int(count); start += groupSize {
			end := start + groupSize
			if end > int(count) {
				end = int(count)
			}

			wg.Add(1)
			go func(start, end int) {
				defer wg.Done()

				for j := start; j < end; j++ {
					// 每次请求一个独立 ctx，避免 defer 在 for 里堆积
					ctx, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)

					height := new(big.Int).Add(startHeight, big.NewInt(int64(j)))

					var h types.Header
					err := c.rpc.CallContext(ctx, &h, "eth_getBlockByNumber", toBlockNumArg(height), false)
					cancel()

					// 只写当前 j 下标，不会越界，也不会与其他 goroutine 冲突
					batchElems[j].Method = "eth_getBlockByNumber"

					if err != nil {
						batchElems[j].Error = err
						batchElems[j].Result = nil
					} else {
						headers[j] = h
						batchElems[j].Result = &headers[j]
						batchElems[j].Error = nil
					}
				}
			}(start, end)
		}

		wg.Wait()
	}

	// ---------- 统一解析阶段：处理错误 / 截断 / 连续性检查 ----------

	size := 0
	var firstErr error

	for i := 0; i < int(count); i++ {
		be := batchElems[i]

		// 1) RPC 层错误处理
		if be.Error != nil {
			if firstErr == nil {
				firstErr = be.Error
			}
			// 如果一条都没拿到就报错 -> 直接 return
			if size == 0 {
				return nil, be.Error
			}
			// 已经拿到部分 header，再遇到错误 -> 截断返回已有部分
			break
		}

		// 2) RPC 层无结果（通常是范围超过最新区块，返回 null）
		if be.Result == nil {
			if size == 0 {
				// 第一条就没有结果，说明这个高度压根不存在
				return nil, fmt.Errorf("no header returned for block %v",
					new(big.Int).Add(startHeight, big.NewInt(int64(i))))
			}
			// 已经拿到部分，再遇到 nil -> 截断即可
			break
		}

		// 3) 类型断言保护，防止 RPC 返回奇怪类型导致 panic
		header, ok := be.Result.(*types.Header)
		if !ok || header == nil {
			if size == 0 {
				return nil, fmt.Errorf("unexpected RPC result type %T at index %d", be.Result, i)
			}
			// 后面就当作链尾截断
			break
		}

		// 4) 写入连续的 headers[size]（注意是 size，而不是 i）
		headers[size] = *header

		// 5) 可选：区块连续性校验（如果你业务不强制要求，可以注释掉）
		if size > 0 {
			if headers[size].ParentHash != headers[size-1].Hash() {
				return nil, fmt.Errorf(
					"non-contiguous headers: header %s does not follow parent %s",
					headers[size].Hash(), headers[size-1].Hash(),
				)
			}
		}

		size++
	}

	if size == 0 {
		// 保险兜底：完全没拿到 header
		if firstErr != nil {
			return nil, firstErr
		}
		return nil, fmt.Errorf("no headers returned in range [%v, %v]", startHeight, endHeight)
	}

	return headers[:size], nil
}

// func (c *clnt) BlockHeadersByRange(startHeight, endHeight *big.Int, chainId uint) ([]types.Header, error) {
// 	// avoid the batch call if there's no range
// 	if startHeight.Cmp(endHeight) == 0 {
// 		header, err := c.BlockHeaderByNumber(startHeight)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return []types.Header{*header}, nil
// 	}

// 	count := new(big.Int).Sub(endHeight, startHeight).Uint64() + 1
// 	headers := make([]types.Header, count)
// 	batchElems := make([]rpc.BatchElem, count)

// 	if chainId != uint(global_const.PolygonChainId) {
// 		for i := uint64(0); i < count; i++ {
// 			height := new(big.Int).Add(startHeight, new(big.Int).SetUint64(i))
// 			batchElems[i] = rpc.BatchElem{Method: "eth_getBlockByNumber", Args: []interface{}{toBlockNumArg(height), false}, Result: &headers[i]}
// 		}

// 		ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
// 		defer cancel()
// 		err := c.rpc.BatchCallContext(ctxwt, batchElems)
// 		if err != nil {
// 			return nil, err
// 		}
// 	} else {
// 		groupSize := 100
// 		var wg sync.WaitGroup
// 		numGroups := (int(count)-1)/groupSize + 1
// 		wg.Add(numGroups)

// 		for i := 0; i < int(count); i += groupSize {
// 			start := i
// 			end := i + groupSize - 1
// 			if end > int(count) {
// 				end = int(count) - 1
// 			}
// 			go func(start, end int) {
// 				defer wg.Done()
// 				for j := start; j <= end; j++ {
// 					ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
// 					defer cancel()
// 					height := new(big.Int).Add(startHeight, new(big.Int).SetUint64(uint64(j)))
// 					batchElems[j] = rpc.BatchElem{
// 						Method: "eth_getBlockByNumber",
// 						Result: new(types.Header),
// 						Error:  nil,
// 					}
// 					header := new(types.Header)
// 					batchElems[j].Error = c.rpc.CallContext(ctxwt, header, batchElems[j].Method, toBlockNumArg(height), false)
// 					batchElems[j].Result = header
// 				}
// 			}(start, end)
// 		}

// 		wg.Wait()
// 	}
// 	// Parse the headers.
// 	//  - Ensure integrity that they build on top of each other
// 	//  - Truncate out headers that do not exist (endHeight > "latest")
// 	size := 0
// 	for i, batchElem := range batchElems {
// 		//if batchElem.Error != nil {
// 		//	if size == 0 {
// 		//		return nil, batchElem.Error
// 		//	} else {
// 		//		break // try return whatever headers are available
// 		//	}
// 		//} else if batchElem.Result == nil {
// 		//	break
// 		//}
// 		header, ok := batchElem.Result.(*types.Header)
// 		if !ok {
// 			return nil, fmt.Errorf("unable to transform rpc response %v into utils.Header", batchElem.Result)
// 		}

// 		headers[i] = *header

// 		//if i > 0 && headers[i].ParentHash != headers[i-1].Hash() {
// 		//	return nil, fmt.Errorf("queried header %s does not follow parent %s", headers[i].Hash(), headers[i-1].Hash())
// 		//}

// 		size = size + 1
// 	}
// 	headers = headers[:size]

// 	return headers, nil
// }

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

	var tx *types.Receipt
	err := c.rpc.CallContext(ctxwt, &tx, "eth_getTransactionReceipt", hash)
	if err != nil {
		return nil, err
	} else if tx == nil {
		return nil, ethereum.NotFound
	}

	return tx, nil
}

// StorageHash returns the sha3 of the storage root for the specified account
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

// FilterLogs returns logs that fit the query parameters. The underlying request is a batch
// request including `eth_getBlockByNumber` to allow the caller to check that connected
// node has the state necessary to fulfill this request
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

	ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout*10)
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

// Modeled off op-service/client.go. We can refactor this once the client/metrics portion
// of op-service/client has been generalized

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

// Needed private utils from geth

func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	if number.Sign() >= 0 {
		return hexutil.EncodeBig(number)
	}
	// It's negative.
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
			// Fail open if we can't figure out what the port should be
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
