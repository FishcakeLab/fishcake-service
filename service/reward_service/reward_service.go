package reward_service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/FishcakeLab/fishcake-service/config"
)

// TokenType represents the type of token for transactions
type TokenType int

const (
	Native TokenType = iota // Native token
	ERC20                   // ERC20 token
)

var (
	StandardTransferGasLimit uint64 = 21000
	ERC20TransferGasLimit    uint64 = 150000
	MaxPriorityFeePerGas            = big.NewInt(30000000000)
)

type RewardService interface {
	DecryptPrivateKey() ([]byte, common.Address, error)
	FccAddress() string
	UsdtAddress() string
	CreateOfflineTransaction(ChainID *big.Int, tokenType TokenType, privateKeyBytes []byte, toAddress string, nonce uint64, gasFeeCap *big.Int, amount *big.Int) (string, string, error)
}

type rewardService struct {
	cfg *config.Config
}

func NewRewardService(cfg *config.Config) RewardService {
	return &rewardService{
		cfg: cfg,
	}
}

func (s *rewardService) DecryptPrivateKey() ([]byte, common.Address, error) {
	ciphertext, err := hex.DecodeString(s.cfg.EncryptedPrivateKey)
	if err != nil {
		log.Error("Decode encrypt private key fail", "err", err)
		return nil, common.Address{}, err
	}
	nonceBytes, err := hex.DecodeString(s.cfg.Nonce)
	if err != nil {
		log.Error("Decode nonce fail", "err", err)
		return nil, common.Address{}, err
	}

	hasher := sha256.New()
	hasher.Write([]byte(s.cfg.KeyPhrase))
	key := hasher.Sum(nil)

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Error("New cipher fail", "err", err)
		return nil, common.Address{}, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Error("New gcm", "err", err)
		return nil, common.Address{}, err
	}

	privateKeyBytes, err := gcm.Open(nil, nonceBytes, ciphertext, nil)
	if err != nil {
		log.Error("Open gcm fail", "err", err)
		return nil, common.Address{}, err
	}

	privateKey, _ := crypto.ToECDSA(privateKeyBytes)

	// 3. Generate address from public key
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	return privateKeyBytes, address, nil
}

func (s *rewardService) FccAddress() string {
	return s.cfg.FCC
}

func (s *rewardService) UsdtAddress() string {
	return s.cfg.USDT
}

func (s *rewardService) CreateOfflineTransaction(ChainID *big.Int, tokenType TokenType, privateKeyBytes []byte, toAddress string, nonce uint64, gasFeeCap *big.Int, amount *big.Int) (string, string, error) {
	switch tokenType {
	case Native:
		return s.createNativeTransaction(ChainID, privateKeyBytes, toAddress, nonce, gasFeeCap, amount)
	case ERC20:
		return s.createERC20Transaction(ChainID, privateKeyBytes, toAddress, nonce, gasFeeCap, amount)
	default:
		return "", "", fmt.Errorf("unsupported token type: %v", tokenType)
	}
}

func (s *rewardService) createNativeTransaction(ChainID *big.Int, privateKeyBytes []byte, toAddress string, nonce uint64, gasFeeCap *big.Int, amount *big.Int) (string, string, error) {
	to := common.HexToAddress(toAddress)
	dFeeTx := &types.DynamicFeeTx{
		ChainID:   ChainID,
		Nonce:     nonce,
		GasTipCap: MaxPriorityFeePerGas,
		GasFeeCap: gasFeeCap,
		Gas:       StandardTransferGasLimit,
		To:        &to,
		Value:     amount,
		Data:      nil,
	}
	privateKeyHex := hex.EncodeToString(privateKeyBytes)
	return OfflineSignTx(dFeeTx, privateKeyHex, ChainID)
}

func (s *rewardService) createERC20Transaction(ChainID *big.Int, privateKeyBytes []byte, toAddress string, nonce uint64, gasFeeCap *big.Int, amount *big.Int) (string, string, error) {
	privateKeyHex := hex.EncodeToString(privateKeyBytes)

	contractAddr := s.cfg.FCC
	to := common.HexToAddress(contractAddr)

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := crypto.Keccak256(transferFnSignature)
	methodID := hash[:4]

	paddedAddress := common.LeftPadBytes(common.HexToAddress(toAddress).Bytes(), 32)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	dFeeTx := &types.DynamicFeeTx{
		ChainID:   ChainID,
		Nonce:     nonce,
		GasTipCap: MaxPriorityFeePerGas,
		GasFeeCap: gasFeeCap,
		Gas:       ERC20TransferGasLimit,
		To:        &to,
		Value:     big.NewInt(0),
		Data:      data,
	}
	return OfflineSignTx(dFeeTx, privateKeyHex, ChainID)
}

func OfflineSignTx(txData *types.DynamicFeeTx, privateKey string, chainId *big.Int) (string, string, error) {
	privateKeyEcdsa, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", "", err
	}
	tx := types.NewTx(txData)
	signer := types.LatestSignerForChainID(chainId)
	signedTx, err := types.SignTx(tx, signer, privateKeyEcdsa)
	if err != nil {
		return "", "", err
	}
	signedTxData, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		return "", "", err
	}
	return "0x" + hex.EncodeToString(signedTxData)[4:], signedTx.Hash().String(), nil
}
