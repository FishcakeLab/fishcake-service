package node

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/log"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/FishcakeLab/fishcake-service/common/bigint"
)

var (
	ErrHeaderTraversalAheadOfProvider            = errors.New("the HeaderTraversal's internal state is ahead of the provider")
	ErrHeaderTraversalAndProviderMismatchedState = errors.New("the HeaderTraversal and provider have diverged in state")
)

type HeaderTraversal struct {
	ethClient EthClient
	chainId   uint

	latestHeader        *types.Header
	lastTraversedHeader *types.Header

	confirmationMode       string
	blockConfirmationDepth *big.Int
	fallbackDepth          *big.Int
}

func NewHeaderTraversal(ethClient EthClient, fromHeader *types.Header, confirmationMode string, confDepth, fallbackDepth *big.Int, chainId uint) *HeaderTraversal {
	if confirmationMode == "" {
		confirmationMode = "depth"
	}
	if confDepth == nil {
		confDepth = bigint.Zero
	}
	if fallbackDepth == nil {
		fallbackDepth = confDepth
	}
	return &HeaderTraversal{
		ethClient:              ethClient,
		lastTraversedHeader:    fromHeader,
		confirmationMode:       strings.ToLower(confirmationMode),
		blockConfirmationDepth: confDepth,
		fallbackDepth:          fallbackDepth,
		chainId:                chainId,
	}
}

func (f *HeaderTraversal) LatestHeader() *types.Header {
	return f.latestHeader
}

func (f *HeaderTraversal) LastTraversedHeader() *types.Header {
	return f.lastTraversedHeader
}

func (f *HeaderTraversal) NextHeaders(maxSize uint64) ([]types.Header, error) {
	latestHeader, err := f.ethClient.BlockHeaderByNumber(nil)
	if err != nil {
		return nil, fmt.Errorf("unable to query latest block: %w", err)
	} else if latestHeader == nil {
		return nil, fmt.Errorf("latest header unreported")
	} else {
		f.latestHeader = latestHeader
	}

	endHeight, err := f.confirmationTargetHeight(latestHeader)
	if err != nil {
		return nil, err
	}
	if endHeight.Sign() < 0 {
		return nil, nil
	}

	if f.lastTraversedHeader != nil {
		cmp := f.lastTraversedHeader.Number.Cmp(endHeight)
		if cmp == 0 {
			return nil, nil
		} else if cmp > 0 {
			return nil, ErrHeaderTraversalAheadOfProvider
		}
	}

	nextHeight := bigint.Zero
	if f.lastTraversedHeader != nil {
		nextHeight = new(big.Int).Add(f.lastTraversedHeader.Number, bigint.One)
	}

	endHeight = bigint.Clamp(nextHeight, endHeight, maxSize)
	headers, err := f.ethClient.BlockHeadersByRange(nextHeight, endHeight, f.chainId)
	if err != nil {
		return nil, fmt.Errorf("error querying blocks by range: %w", err)
	}

	numHeaders := len(headers)
	if numHeaders == 0 {
		return nil, nil
	} else if f.lastTraversedHeader != nil && headers[0].ParentHash != f.lastTraversedHeader.Hash() {
		log.Error("block hash is not match", "lastTraversedHeader", f.lastTraversedHeader.Number, "headers[0]", headers[0].Number)
		// return nil, ErrHeaderTraversalAndProviderMismatchedState
	}

	f.lastTraversedHeader = &headers[numHeaders-1]
	return headers, nil
}

func (f *HeaderTraversal) confirmationTargetHeight(latestHeader *types.Header) (*big.Int, error) {
	switch f.confirmationMode {
	case "finalized":
		finalizedHeader, err := f.ethClient.LatestFinalizedBlockHeader()
		if err == nil && finalizedHeader != nil && finalizedHeader.Number != nil {
			return new(big.Int).Set(finalizedHeader.Number), nil
		}
		log.Warn("failed to fetch finalized header, falling back to confirmation depth", "err", err, "fallbackDepth", f.fallbackDepth)
		return confirmedHeight(latestHeader.Number, f.fallbackDepth), nil
	case "safe":
		safeHeader, err := f.ethClient.LatestSafeBlockHeader()
		if err == nil && safeHeader != nil && safeHeader.Number != nil {
			return new(big.Int).Set(safeHeader.Number), nil
		}
		log.Warn("failed to fetch safe header, falling back to confirmation depth", "err", err, "fallbackDepth", f.fallbackDepth)
		return confirmedHeight(latestHeader.Number, f.fallbackDepth), nil
	case "depth", "":
		return confirmedHeight(latestHeader.Number, f.blockConfirmationDepth), nil
	default:
		log.Warn("unknown confirmation mode, falling back to confirmation depth", "mode", f.confirmationMode, "confirmationDepth", f.blockConfirmationDepth)
		return confirmedHeight(latestHeader.Number, f.blockConfirmationDepth), nil
	}
}

func confirmedHeight(latestNumber, depth *big.Int) *big.Int {
	if latestNumber == nil {
		return nil
	}
	if depth == nil {
		return new(big.Int).Set(latestNumber)
	}
	return new(big.Int).Sub(latestNumber, depth)
}
